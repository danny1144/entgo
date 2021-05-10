package main

import (
	"context"
	"ego/ent"
	"ego/ent/car"
	"ego/ent/group"
	"ego/ent/migrate"
	"ego/ent/user"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	client, err := ent.Open("mysql", "api:123456@tcp(localhost:3306)/api?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	//可以不生成外键约束
	if err := client.Schema.Create(context.Background(),migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources :%v", err)
	}
	//创建用户
	//CreateUser(context.Background(), client)
	//查询用户
	//u, _ := QueryUser(context.Background(), client)
	// 创建用户并添加2两车
	//CreateCar(context.Background(), client)

	//QueryCars(context.Background(), u)

	//创建图
	//CreateGraph(context.Background(), client)
	QueryGithub(context.Background(), client)
}
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().
		SetName("张三").SetAge(19).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user :%w", err)
	}

	return u, nil
}
func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Query().Where(user.NameEQ("zs")).First(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user :%w", err)
	}

	log.Println("return user:", u)
	return u, nil
}

func CreateCar(ctx context.Context, client *ent.Client) (*ent.User, error) {
	tesla, err := client.Car.Create().
		SetModel("Tesla").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car:%w", err)
	}
	log.Println("car was created", tesla)

	audi, err := client.Car.Create().
		SetModel("audi").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car:%w", err)
	}
	log.Println("car was created", audi)

	zs, err := client.User.Create().
		SetAge(18).SetName("zs").AddCars(tesla, audi).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("user was created : ", zs)

	return zs, nil

}
func QueryCars(ctx context.Context, user *ent.User) error {

	cars, err := user.QueryCars().All(ctx)

	if err != nil {
		return fmt.Errorf("failed querying user cars:%w", err)
	}

	log.Println("return cars :", cars)

	audi, err := user.QueryCars().Where(car.ModelEQ("audi")).Only(ctx)

	if err != nil {
		return fmt.Errorf("failed querying user cars:%w", err)
	}

	log.Println(audi)
	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// 首先创建一些用户
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}
	// 然后，创建一些汽车，并在创建时就关联到用户
	_, err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m). // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m). // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(neta). // attach this graph to Neta.
		Save(ctx)
	if err != nil {
		return err
	}
	// 创建用户组，并在创建之初就为他们添加一些用户
	_, err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}
func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.
		Query().
		Where(group.Name("GitHub")). // (Group(Name=GitHub),)
		QueryUsers(). // (User(Name=Ariel, Age=30),)
		QueryCars(). // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
	return nil
}
