// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package model

import "database/sql"

func ScanUser(r *sql.Row) (User, error) {
	var s User
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Email,
		&s.Salt,
		&s.Salted,
		&s.Created,
		&s.Updated,
	); err != nil {
		return User{}, err
	}
	return s, nil
}

func ScanUsers(rs *sql.Rows) ([]User, error) {
	structs := make([]User, 0, 16)
	var err error
	for rs.Next() {
		var s User
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Email,
			&s.Salt,
			&s.Salted,
			&s.Created,
			&s.Updated,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanComment(r *sql.Row) (Comment, error) {
	var s Comment
	if err := r.Scan(
		&s.Id,
		&s.Message,
		&s.Created,
		&s.Updated,
	); err != nil {
		return Comment{}, err
	}
	return s, nil
}

func ScanComments(rs *sql.Rows) ([]Comment, error) {
	structs := make([]Comment, 0, 16)
	var err error
	for rs.Next() {
		var s Comment
		if err = rs.Scan(
			&s.Id,
			&s.Message,
			&s.Created,
			&s.Updated,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

