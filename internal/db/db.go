// Code generated by sqlc. DO NOT EDIT.

package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addProductStmt, err = db.PrepareContext(ctx, addProduct); err != nil {
		return nil, fmt.Errorf("error preparing query AddProduct: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteProductStmt, err = db.PrepareContext(ctx, deleteProduct); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProduct: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getAllMerchantProductsStmt, err = db.PrepareContext(ctx, getAllMerchantProducts); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllMerchantProducts: %w", err)
	}
	if q.getMerchantIDStmt, err = db.PrepareContext(ctx, getMerchantID); err != nil {
		return nil, fmt.Errorf("error preparing query GetMerchantID: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserPasswordStmt, err = db.PrepareContext(ctx, getUserPassword); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserPassword: %w", err)
	}
	if q.getUsersStmt, err = db.PrepareContext(ctx, getUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsers: %w", err)
	}
	if q.insertMerchantProductStmt, err = db.PrepareContext(ctx, insertMerchantProduct); err != nil {
		return nil, fmt.Errorf("error preparing query InsertMerchantProduct: %w", err)
	}
	if q.isUsernameExistStmt, err = db.PrepareContext(ctx, isUsernameExist); err != nil {
		return nil, fmt.Errorf("error preparing query IsUsernameExist: %w", err)
	}
	if q.updateUserPasswordStmt, err = db.PrepareContext(ctx, updateUserPassword); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserPassword: %w", err)
	}
	if q.updateUserToMerchantStmt, err = db.PrepareContext(ctx, updateUserToMerchant); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserToMerchant: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addProductStmt != nil {
		if cerr := q.addProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addProductStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteProductStmt != nil {
		if cerr := q.deleteProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProductStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getAllMerchantProductsStmt != nil {
		if cerr := q.getAllMerchantProductsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllMerchantProductsStmt: %w", cerr)
		}
	}
	if q.getMerchantIDStmt != nil {
		if cerr := q.getMerchantIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMerchantIDStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserPasswordStmt != nil {
		if cerr := q.getUserPasswordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserPasswordStmt: %w", cerr)
		}
	}
	if q.getUsersStmt != nil {
		if cerr := q.getUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersStmt: %w", cerr)
		}
	}
	if q.insertMerchantProductStmt != nil {
		if cerr := q.insertMerchantProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertMerchantProductStmt: %w", cerr)
		}
	}
	if q.isUsernameExistStmt != nil {
		if cerr := q.isUsernameExistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing isUsernameExistStmt: %w", cerr)
		}
	}
	if q.updateUserPasswordStmt != nil {
		if cerr := q.updateUserPasswordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserPasswordStmt: %w", cerr)
		}
	}
	if q.updateUserToMerchantStmt != nil {
		if cerr := q.updateUserToMerchantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserToMerchantStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                         DBTX
	tx                         *sql.Tx
	addProductStmt             *sql.Stmt
	createUserStmt             *sql.Stmt
	deleteProductStmt          *sql.Stmt
	deleteUserStmt             *sql.Stmt
	getAllMerchantProductsStmt *sql.Stmt
	getMerchantIDStmt          *sql.Stmt
	getUserStmt                *sql.Stmt
	getUserPasswordStmt        *sql.Stmt
	getUsersStmt               *sql.Stmt
	insertMerchantProductStmt  *sql.Stmt
	isUsernameExistStmt        *sql.Stmt
	updateUserPasswordStmt     *sql.Stmt
	updateUserToMerchantStmt   *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                         tx,
		tx:                         tx,
		addProductStmt:             q.addProductStmt,
		createUserStmt:             q.createUserStmt,
		deleteProductStmt:          q.deleteProductStmt,
		deleteUserStmt:             q.deleteUserStmt,
		getAllMerchantProductsStmt: q.getAllMerchantProductsStmt,
		getMerchantIDStmt:          q.getMerchantIDStmt,
		getUserStmt:                q.getUserStmt,
		getUserPasswordStmt:        q.getUserPasswordStmt,
		getUsersStmt:               q.getUsersStmt,
		insertMerchantProductStmt:  q.insertMerchantProductStmt,
		isUsernameExistStmt:        q.isUsernameExistStmt,
		updateUserPasswordStmt:     q.updateUserPasswordStmt,
		updateUserToMerchantStmt:   q.updateUserToMerchantStmt,
	}
}
