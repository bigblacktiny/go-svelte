package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"goTemp/globalUtils"
	"goTemp/globalerrors"
	pb "goTemp/user/proto"
	"goTemp/user/server/statements"
	"log"
	"time"
)

//glErr: Holds the service global errors that are shared cross services
var glErr globalerrors.SrvError

//Promotion: Main entry point for promotion related services
type User struct{}

//promoErr: Holds service specific errors
var userErr statements.UserErr

func (u *User) GetUserById(ctx context.Context, searchId *pb.SearchId, outUser *pb.User) error {
	_ = ctx

	var validFrom time.Time
	var validThru time.Time

	sqlStatement := statements.SqlSelectById.String()
	err := conn.QueryRow(context.Background(), sqlStatement,
		searchId.GetId()).
		Scan(
			&outUser.Id,
			&outUser.Firstname,
			&outUser.Lastname,
			&validFrom,
			&validThru,
			&outUser.Active,
			&outUser.Pwd,
			&outUser.Name,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		} else {
			log.Printf(userErr.SelectRowReadError(err))
			return err
		}

	}

	convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru)
	if err != nil {
		return err
	}
	outUser.ValidFrom, outUser.ValidThru = convertedTimes[0], convertedTimes[1]

	return nil
}

func (u *User) GetUsers(ctx context.Context, searchParms *pb.SearchParams, users *pb.Users) error {

	_ = ctx

	sqlStatement := statements.SqlSelectAll.String()
	sqlWhereClause, values, err2 := u.buildSearchWhereClause(searchParms)
	if err2 != nil {
		return err2
	}

	sqlStatement += sqlWhereClause

	rows, err := conn.Query(context.Background(), sqlStatement, values...)

	if err != nil {
		log.Printf(userErr.SelectReadError(err))
		return err
	}

	var validFrom time.Time
	var validThru time.Time
	for rows.Next() {
		var user pb.User
		err := rows.
			Scan(
				&user.Id,
				&user.Firstname,
				&user.Lastname,
				&validFrom,
				&validThru,
				&user.Active,
				&user.Pwd,
				&user.Name,
			)
		if err != nil {
			log.Printf(userErr.SelectScanError(err))
			return err
		}

		convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru)
		if err != nil {
			return err
		}
		user.ValidFrom, user.ValidThru = convertedTimes[0], convertedTimes[1]

		users.User = append(users.User, &user)
	}

	return nil
}

//buildSearchWhereClause: Builds a sql string to be used as the where clause in a sql statement. It also returns an interface
//slice with the values to be used as replacements in the sql statement. Currently only handles equality constraints, except
//for the date lookup which is done  as a contains clause
func (u *User) buildSearchWhereClause(searchParms *pb.SearchParams) (string, []interface{}, error) {
	sqlWhereClause := " where 1=1"
	var values []interface{}

	i := 1
	if searchParms.GetId() != 0 {
		sqlWhereClause += fmt.Sprintf(" AND user.id = $%d", i)
		values = append(values, searchParms.GetId())
		i++
	}
	if searchParms.GetFisrtname() != "" {
		sqlWhereClause += fmt.Sprintf(" AND user.firstname = $%d", i)
		values = append(values, searchParms.GetFisrtname())
		i++
	}
	if searchParms.GetLastname() != "" {
		sqlWhereClause += fmt.Sprintf(" AND user.lastname = $%d", i)
		values = append(values, searchParms.GetLastname())
		i++
	}
	if searchParms.GetValidDate() != nil {
		convertedDates, err := globalUtils.TimeStampPPBToTime(searchParms.GetValidDate())
		if err != nil {
			return "", nil, err
		}
		validFrom := convertedDates[0]
		sqlWhereClause += fmt.Sprintf(" AND user.validfrom <= $%d AND user.validthru >= $%d", i, i)
		values = append(values, validFrom)
		i++
	}
	return sqlWhereClause, values, nil
}

func (u *User) CreateUser(ctx context.Context, inUser *pb.User, outUser *pb.User) error {
	_ = ctx

	if errVal := u.BeforeCreateUser(ctx, inUser, &pb.ValidationErr{}); errVal != nil {
		return errVal
	}

	convertedDates, err := globalUtils.TimeStampPPBToTime(inUser.GetValidFrom(), inUser.GetValidThru())
	if err != nil {
		return err
	}
	validFrom, validThru := convertedDates[0], convertedDates[1]

	//id, firstname , lastname, validfrom, validthru, active, pwd, name

	sqlStatement := statements.SqlInsert.String()
	errIns := conn.QueryRow(context.Background(), sqlStatement,
		inUser.GetFirstname(),
		inUser.GetLastname(),
		validFrom,
		validThru,
		inUser.GetActive(),
		inUser.GetPwd(),
		inUser.GetName(),
	).
		Scan(
			&outUser.Id,
			&outUser.Firstname,
			&outUser.Lastname,
			&validFrom,
			&validThru,
			&outUser.Active,
			&outUser.Pwd,
			&outUser.Name,
		)

	if errIns != nil {
		log.Printf(userErr.InsertError(err))
		return errIns
	}

	convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru)
	if err != nil {
		return err
	}
	outUser.ValidFrom, outUser.ValidThru = convertedTimes[0], convertedTimes[1]

	if errVal := u.AfterCreateUser(ctx, outUser, &pb.AfterFuncErr{}); errVal != nil {
		return errVal
	}

	return nil
}

func (u *User) UpdateUser(ctx context.Context, inUser *pb.User, outUser *pb.User) error {
	_ = ctx

	if errVal := u.BeforeUpdateUser(ctx, inUser, &pb.ValidationErr{}); errVal != nil {
		return errVal
	}

	sqlStatement := statements.SqlUpdate.String()

	convertedDates, err := globalUtils.TimeStampPPBToTime(inUser.GetValidFrom(), inUser.GetValidThru())
	if err != nil {
		return err
	}
	validFrom, validThru := convertedDates[0], convertedDates[1]

	err = conn.QueryRow(context.Background(), sqlStatement,
		inUser.GetFirstname(),
		inUser.GetLastname(),
		validFrom,
		validThru,
		inUser.GetActive(),
		inUser.GetPwd(),
		inUser.GetName(),
		inUser.GetId(),
	).Scan(
		&outUser.Id,
		&outUser.Firstname,
		&outUser.Lastname,
		&validFrom,
		&validThru,
		&outUser.Active,
		&outUser.Pwd,
		&outUser.Name,
	)
	if err != nil {
		log.Printf(userErr.UpdateError(err))
		return err
	}

	convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru)
	if err != nil {
		return err
	}
	outUser.ValidFrom, outUser.ValidThru = convertedTimes[0], convertedTimes[1]

	if errVal := u.AfterUpdateUser(ctx, outUser, &pb.AfterFuncErr{}); errVal != nil {
		return errVal
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, searchid *pb.SearchId, affectedCount *pb.AffectedCount) error {
	_ = ctx

	outUser := pb.User{}
	if err := u.GetUserById(ctx, searchid, &outUser); err != nil {
		return err
	}
	if errVal := u.BeforeDeleteUser(ctx, &outUser, &pb.ValidationErr{}); errVal != nil {
		return errVal
	}

	sqlStatement := statements.SqlDelete.String()
	commandTag, err := conn.Exec(context.Background(), sqlStatement, searchid.Id)
	if err != nil {
		log.Printf(userErr.DeleteError(searchid.Id, err))
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf(userErr.DeleteRowNotFoundError(searchid.Id))
	}

	affectedCount.Value = commandTag.RowsAffected()

	if errVal := u.AfterDeleteUser(ctx, &outUser, &pb.AfterFuncErr{}); errVal != nil {
		return errVal
	}

	return nil
}