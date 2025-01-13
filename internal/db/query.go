package db

const TotalQuery = `
	select
		count(*) total_items,	
		count(distinct category) total_categories, 
		sum(price) total_price
	from prices p
`
