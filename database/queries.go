package database

const SQL_GET_QUERY = `
	select
		p.product_id,
		p.name,
		p.price,
		c.category_id,
		c.name
	from products as p
	join categories as c using(category_id);
`

const SQL_POST_QUERY = `
	insert into products(name, price, category_id) values($1, $2, $3)
	returning
		product_id,
		name,
		price,
		category_id;
`

const SQL_DELETE_QUERY = `
	delete from products where product_id = $1	
	returning
		product_id,
		name,
		price,
		category_id;
`

const SQL_PUT_QUERY = `
	update products
	set
		name = case when $2 = '' then name else $2 end,
		price = case when $3 = 0 then price else $3 end,
		category_id = case when $4 = 0 then category_id else $4 end
	where product_id = $1
	returning
		product_id,
		name, 
		price,
		category_id;
`
