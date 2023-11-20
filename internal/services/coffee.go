package services

import (
	"context"
	"time"
)

type Coffee struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Roast string `json:"roast"`
	Image string `json:"image"`
	Region string `json:"region"`
	Price float32 `json:"price"`
	GrindUnit int16 `json:"grind_unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Coffee) GetAllCofees() ([]*Coffee, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, name, roast, image, region, price, grind_unit, created_at, updated_at FROM coffees`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var coffees []*Coffee
	for rows.Next(){
		var coffee Coffee
		err := rows.Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Roast,
			&coffee.Image,
			&coffee.Region,
			&coffee.Price,
			&coffee.GrindUnit,
			&coffee.CreatedAt,
			&coffee.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		coffees = append(coffees, &coffee)
	}
	return coffees, nil
}

func (c *Coffee) GetCoffeeById(id string) (*Coffee, error){

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, name, roast, image, region, price, grind_unit, created_at, updated_at FROM coffees WHERE id = $1` 
	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&c.ID,
		&c.Name,
		&c.Roast,
		&c.Image,
		&c.Region,
		&c.Price,
		&c.GrindUnit,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Coffee) CreateCoffee(coffee Coffee) (*Coffee, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `INSERT INTO coffees (name, roast, image, region, price, grind_unit, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at, updated_at`

	_, err := db.ExecContext(
		ctx,
		query,
		coffee.Name,
		coffee.Roast,
		coffee.Image,
		coffee.Region,
		coffee.Price,
		coffee.GrindUnit,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return &coffee, nil

}