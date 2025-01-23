package db

const TotalQuery = `
	select	
		count(distinct category) total_categories, 
		sum(price) total_price
	from prices p
`

const GetItems = `
	select
		id,
		create_date,
		name,
		category,
		price
	from prices p
`

const AddItem = `
	insert into prices (create_date, name, category, price) values ($1, $2, $3, $4)
`
