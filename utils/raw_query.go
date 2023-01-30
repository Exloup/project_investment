package utils

// Customer Section
// List all account
const (
	VIEW_ACCOUNTS = `SELECT * FROM mst_account`
)

// Create New Account
const (
	CREATE_ACCOUNTS = `INSERT INTO mst_account(email,password,name,phone_number) VALUES ($1,$2,$3,$4)`
)

// Select Account with ID
const (
	SELECT_ACCOUNT_IDS = `SELECT * FROM mst_account WHERE id =$1`
)

// Update Account with ID
const (
	UPDATE_ACCOUNT_IDS = `UPDATE mst_account set name=$1, phone_number=$2 WHERE id=$3 `
)

// Activated & Deactivated Account with ID
const (
	ACTIVE_ACCOUNT_IDS   = `UPDATE mst_account set is_activated = TRUE WHERE id=$1 `
	DEACTIVE_ACCOUNT_IDS = `UPDATE mst_account set is_activated = FALSE WHERE id=$1`
)

// Delete Account with ID
const (
	DELETE_ACCOUNT = `DELETE FROM mst_account WHERE id=$1`
)

// Product Section
// List all Product
const (
	VIEW_PRODUCT = `SELECT * FROM mst_product`
)

// Create New Product
const (
	INSERT_PRODUCT = `INSERT INTO mst_product(product_name,description_product) VALUES ($1,$2)`
)

// Select Product with ID
const (
	SELECT_PRODUCT_IDS = `SELECT * FROM mst_product WHERE id =$1`
)

// Update Product with ID
const (
	UPDATE_PRODUCT_IDS = `UPDATE mst_product set product_name=$1,description_product=$2 WHERE id=$3 `
)

// Delete Product with ID
const (
	DELETE_PRODUCT = `DELETE FROM mst_product WHERE id=$1`
)

// Portofolio Section
// List all portofolio
const (
	VIEW_PORTOFOLIO = `SELECT * FROM mst_portofolio`
)

// Create New Portofolio
const (
	INSERT_PORTOFOLIO = `INSERT INTO mst_portofolio(account_id,total_quantity,profit) VALUES ($1,$2,$3)`
)

// Select Portofolio with ID
const (
	SELECT_PORTOFOLIO_IDS = `SELECT * FROM mst_portofolio WHERE id =$1`
)

// Update Portofolio with ID
const (
	UPDATE_PORTOFOLIO_IDS = `UPDATE mst_portofolio set account_id=$1,total_quantity=$2 ,profit=$3 WHERE id=$4 `
)

// Delete Portofolio with ID
const (
	DELETE_PORTOFOLIO = `DELETE FROM mst_portofolio WHERE id=$1`
)

// Price Section
// List all Price
const (
	VIEW_PRICE = `SELECT * FROM prices`
)

// Create New Price
const (
	INSERT_PRICE = `INSERT INTO prices(price,date,product_id,performa) VALUES ($1,$2,$3,$4)`
)

// Select Price with ID
const (
	SELECT_PRICE_IDS = `SELECT * FROM prices WHERE id =$1`
)

// Update Price with ID
const (
	UPDATE_PRICE_IDS = `UPDATE prices set price=$1,performa=$2 WHERE id=$3 `
)

// Delete Price with ID
const (
	DELETE_PRICE = `DELETE FROM prices WHERE id=$1`
)

// Transaction Section
// List all Transaction
const (
	VIEW_TRX_PRODUCT = `SELECT * FROM trx_product`
)

// Create New Transaction
const (
	INSERT_TRX_PRODUCT = `INSERT INTO trx_product(trx_code,date,portofolio_id,product_id,status_description) VALUES ($1,$2,$3,$4,$5)`
)

// Select Transaction with ID
const (
	SELECT_TRX_PRODUCT_IDS = `SELECT * FROM trx_product WHERE id =$1`
)

// Update Transaction with ID
const (
	UPDATE_TRX_PRODUCT_IDS = `UPDATE trx_product set trx_code=$1,portofolio_id=$2,product_id=$3,status_description=$4 WHERE id=$5 `
)

// Delete Transaction with ID
const (
	DELETE_TRX_PRODUCT = `DELETE FROM trx_product WHERE id=$1`
)

// Transaction Detail Section
// List all Transaction Detail
const (
	VIEW_TRX_DETAIL = `SELECT * FROM trx_detail`
)

// Create New Transaction Detail
const (
	INSERT_TRX_DETAIL = `INSERT INTO trx_detail(id,price_id,trx_product_id,quantity) VALUES ($1,$2,$3,$4)`
)

// Select Transaction Detail with ID
const (
	SELECT_TRX_DETAIL_IDS = `SELECT * FROM trx_detail WHERE id =$1`
)

// Update Transaction Detail with ID
const (
	UPDATE_TRX_DETAIL_IDS = `UPDATE trx_detail set price_id=$1,trx_product_id=$2,quantity=$3 WHERE id=$4 `
)

// Delete Transaction Detail with ID
const (
	DELETE_TRX_DETAIL = `DELETE FROM trx_detail WHERE id=$1`
)
