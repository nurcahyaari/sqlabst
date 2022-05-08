# Table of Content
<ol>
    <li>
        <a href="">Name</a>
    </li>
    <li>
        <a href="">Getting started</a>
        <ul>
            <a href="">Installation</a>
        </ul>
    </li>
    <li>
        <a href="">Usage</a>
    </li>
</ol>

# SQLABST
sqlabst is the acronym for SQL Abstraction, this is a simple sql abstraction to join sqlx.Tx and sqlx.DB in the same interface. sqlabst also supports filtering and updating builder


# Getting started
This is an example to how to use this project locally

## Installation
    go get github.com/nurcahyaari/sqlabst

# Usage

sqlabst is an abstraction from sqlx, so all of the sqlx's function is supported on sqlabst. the difference is only how sqlabst treats a transaction. because sqlabst combines sqlx.DB and sqlx.Tx, when you call a function that implements on sqlx.DB and sqlx.Tx it's by default called the sqlx.DB's function. but if you start the transaction the function will call its sqlx.Tx own

example:
```go

// start transaction

type Product struct {
	ProductId int64  `db:"product_id"`
	Name      string `db:"name"`
}

type ProductList []*Product

func main() {
    log.Info().Msg("Initialize Mysql connection")

	dbHost := ""
	dbPort := ""
	dbName := ""
	dbUser := ""
	dbPass := ""

	sHost := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("mysql", sHost)

	if err != nil {
		log.Err(err).Msgf("Error to loading Database %s", err)
		panic(err)
	}

	log.Info().Str("Name", dbName).Msg("Success connect to DB")

    sqlabst := sqlabst.NewSqlAbst(db)

    // this query will be fetch on sqlx.DB
    var productList ProductList
	query := "SELECT product_id, name FROM products WHERE product_id IN (?)"

	query, args, err := sqlx.In(query, id)
	if err != nil {
		return
	}

	err = sqlabst.SelectContext(ctx,
		&productList,
		query,
		args...)

    query := "INSERT INTO products (name) VALUES (?)"
	sqlabst.ExecContext(ctx, query, product.Name)


    //this query will be fetch on sqlx.Tx
    sqlabst.Beginx() // when transactions is started by default all of the functions that implement from sqlx.DB and sqlx.Tx will call the sqlx.Tx

    var productList ProductList
	query := "SELECT product_id, name FROM products WHERE product_id IN (?)"

	query, args, err := sqlx.In(query, id)
	if err != nil {
        sqlabst.Rollback()
		return
	}

	err = sqlabst.SelectContext(ctx,
		&productList,
		query,
		args...)
    if err != nil {
        sqlabst.Rollback()
		return
	}

    query := "INSERT INTO products (name) VALUES (?)"
	_, err = sqlabst.ExecContext(ctx, query, product.Name)
    if err != nil {
        sqlabst.Rollback()
        return
	}
}


```