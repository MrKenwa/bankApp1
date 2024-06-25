package runtime

import (
	models2 "bankApp1/internal/models"
	"bankApp1/internal/users/repo/postgres"
	"bankApp1/internal/users/usecase"
	"bankApp1/repo/balanceRepo"
	"bankApp1/repo/cardRepo"
	"bankApp1/repo/depositRepo"
	"bankApp1/repo/operationRepo"
	"bankApp1/txManager"
	"bankApp1/usecase/paymentUsecase"
	"bankApp1/usecase/productsUsecase"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type runtime struct {
	userRep      *postgres.UserRepo
	cardRep      *cardRepo.CardRepo
	depositRep   *depositRepo.DepositRepo
	balanceRep   *balanceRepo.BalanceRepo
	operationRep *operationRepo.OperationRepo
	userUC       *usecase.UserUC
	productUC    *productsUsecase.ProductsUsecase
	paymentUC    *paymentUsecase.PaymentUC
}

func newRuntime(manager *txManager.TxManager) runtime {
	userRep := postgres.NewUserRepo(manager)
	cardRep := cardRepo.NewCardRepo(manager)
	depRepo := depositRepo.NewDepositRepo(manager)
	balRepo := balanceRepo.NewBalanceRepo(manager)
	opRepo := operationRepo.NewOperationRepo(manager)
	userUC := usecase.NewUserUC(manager, userRep)
	productsUC := productsUsecase.NewProductsUsecase(manager, cardRep, depRepo, balRepo)
	payUC := paymentUsecase.NewPaymentUC(manager, balRepo, opRepo)
	return runtime{
		userRep:      userRep,
		cardRep:      cardRep,
		depositRep:   depRepo,
		balanceRep:   balRepo,
		operationRep: opRepo,
		userUC:       userUC,
		productUC:    productsUC,
		paymentUC:    payUC,
	}
}

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

const RegisterUser = "1"
const LoginUser = "2"
const (
	Send         = "1"
	PayIn        = "2"
	PayOut       = "3"
	OpenNew      = "4"
	CloseProduct = "5"
	Exit         = "6"
)

func Run(manager *txManager.TxManager) {
	runtime := newRuntime(manager)
	var uid models2.UserID

	clearConsole()
	fmt.Printf("Welcome to Z bank!\n")
	userIn := 0
	for 1 > userIn {
		fmt.Println("What you will choose?\n1.Register\n2.Log in")
		var ch string
		fmt.Scanln(&ch)
		switch ch {
		case RegisterUser:
			clearConsole()
			user := GetUser()
			if id, err := runtime.userUC.Register(&user); err != nil {
				log.Fatalf("Register failed: %v", err)
			} else {
				userIn = 1
				uid = id
				log.Printf("Register success, your id: %v", uid)
			}
		case LoginUser:
			clearConsole()
			fmt.Println("Enter your email")
			var email string
			fmt.Scanln(&email)
			filter := models2.UserFilter{Emails: []string{email}}
			fmt.Println("Enter your password")
			var password string
			fmt.Scanln(&password)
			if id, err := runtime.userUC.Login(filter, password); err != nil {
				log.Fatalf("Login failed: %v", err)
			} else {
				userIn = 1
				uid = id
				log.Printf("Login success")
			}
		}
	}

	for 2 > userIn {
		clearConsole()
		fmt.Println("What you want to do?\n1.Send money\n2.Pay in balance\n3.Pay out\n4.Open new product\n5.Close product\n6.Exit")
		var ch string
		fmt.Scanln(&ch)
		switch ch {
		case Send:
			clearConsole()
			err := SendMoney(uid, runtime)
			if err != nil {
				log.Fatalf("Send failed: %v", err)
			}
		case PayIn:
			clearConsole()
			err := PayInFunc(runtime)
			if err != nil {
				log.Fatalf("Pay in failed: %v", err)
			}
		case PayOut:
			clearConsole()
			err := PayOutFunc(uid, runtime)
			if err != nil {
				log.Fatalf("Pay out failed: %v", err)
			}
		case OpenNew:
			clearConsole()
			err := NewProductFunc(uid, runtime)
			if err != nil {
				log.Fatalf("Open new product failed: %v", err)
			}
		case CloseProduct:
			clearConsole()
			err := CloseProductFunc(uid, runtime)
			if err != nil {
				log.Fatalf("Close product failed: %v", err)
			}
		case Exit:
			userIn = 2
			clearConsole()
			fmt.Println("Goodbye!")
		}
	}
	return
}

func GetUser() models2.User {
	var user models2.User
	fmt.Println("Enter your name: ")
	fmt.Scanln(&user.Name)
	fmt.Println("Enter your last name")
	fmt.Scanln(&user.Lastname)
	fmt.Println("Enter your patronymic")
	fmt.Scanln(&user.Patronymic)

	fmt.Println("Enter your email: ")
	fmt.Scanln(&user.Email)
	fmt.Println("Enter your password: ")
	fmt.Scanln(&user.Password)
	fmt.Println("Enter your passport number: ")
	fmt.Scanln(&user.PassportNumber)
	user.CreatedAt = time.Now()

	return user
}

func SendMoney(uid models2.UserID, runtime runtime) error {
	senderFilter, err := chooseProduct(uid, runtime)
	if err != nil {
		return err
	}
	clearConsole()

	receiverFilter, err := chooseReceiver()
	clearConsole()

	var amount int
	var ch string
	fmt.Println("Enter amount: ")
	fmt.Scanln(&ch)
	amount, _ = strconv.Atoi(ch)
	_, err = runtime.paymentUC.Send(senderFilter, receiverFilter, int64(amount), "transfer")
	if err != nil {
		return err
	}
	clearConsole()
	fmt.Println("The operation was successful!\nEnter anything to continue...")
	fmt.Scanln()
	return nil
}

func PayOutFunc(uid models2.UserID, runtime runtime) error {
	senderFilter, err := chooseProduct(uid, runtime)
	if err != nil {
		return err
	}

	clearConsole()
	var ch string
	var amount int
	fmt.Println("Enter amount: ")
	fmt.Scanln(&ch)
	amount, _ = strconv.Atoi(ch)
	_, err = runtime.paymentUC.PayOut(senderFilter, int64(amount), "pay out")
	if err != nil {
		return err
	}
	clearConsole()
	fmt.Println("The operation was successful!\nEnter anything to continue...")
	fmt.Scanln()
	return nil
}

func PayInFunc(runtime runtime) error {
	receiverFilter, err := chooseReceiver()
	if err != nil {
		return err
	}
	clearConsole()
	var ch string
	var amount int
	fmt.Println("Enter amount: ")
	fmt.Scanln(&ch)
	amount, _ = strconv.Atoi(ch)
	_, err = runtime.paymentUC.PayIn(receiverFilter, int64(amount), "pay in")
	if err != nil {
		return err
	}
	clearConsole()
	fmt.Println("The operation was successful!\nEnter anything to continue...")
	fmt.Scanln()
	return nil
}

func NewProductFunc(uid models2.UserID, runtime runtime) error {
	fmt.Println("Choose type: ")
	fmt.Println("1.Card\n2.Deposit")
	var ch string
	fmt.Scanln(&ch)
	switch ch {
	case "1":
		clearConsole()
		var pin string
		fmt.Println("Enter pin: ")
		fmt.Scanln(&pin)
		cid, err := runtime.productUC.CreateNewCard(uid, "debit card", pin)
		if err != nil {
			return err
		}
		clearConsole()
		fmt.Println("New card id:", cid)
	case "2":
		clearConsole()
		did, err := runtime.productUC.CreateNewDeposit(uid, "deposit", 1.25)
		if err != nil {
			return err
		}
		clearConsole()
		fmt.Println("New deposit id:", did)
	}

	fmt.Println("Enter anything to continue...")
	fmt.Scanln()
	return nil
}

func CloseProductFunc(uid models2.UserID, runtime runtime) error {
	productFilter, err := chooseProduct(uid, runtime)
	if err != nil {
		return err
	}
	if len(productFilter.CardIDs) != 0 {
		err = runtime.productUC.DeleteCard(productFilter.CardIDs[0])
		if err != nil {
			return err
		}
	} else if len(productFilter.DepositIDs) != 0 {
		err = runtime.productUC.DeleteDeposit(productFilter.DepositIDs[0])
		if err != nil {
			return err
		}
	}
	fmt.Println("Successful delete product!")
	fmt.Println("Enter anything to continue...")
	fmt.Scanln()
	return nil
}

func chooseProduct(uid models2.UserID, runtime runtime) (models2.BalanceFilter, error) {
	fmt.Println("Choose your product:")
	cards, err := runtime.productUC.GetCards(uid)
	if err != nil {
		fmt.Printf("Get cards failed: %v", err)
	}
	deposits, err := runtime.productUC.GetDeposits(uid)
	if err != nil {
		fmt.Printf("Get deposits failed: %v", err)
	}

	count := 1
	for _, card := range cards {
		fmt.Printf("%d.%s id: %d\n", count, card.Type, card.CardID)
		count++
	}

	for _, deposit := range deposits {
		fmt.Printf("%d.%s id: %d\n", count, deposit.Type, deposit.DepositID)
		count++
	}

	var senderFilter models2.BalanceFilter
	var ch string
	fmt.Scanln(&ch)
	n, _ := strconv.Atoi(ch)
	n--
	if n < len(cards) {
		senderFilter.CardIDs = []models2.CardID{cards[n].CardID}
	} else if n < len(cards)+len(deposits) {
		senderFilter.DepositIDs = []models2.DepositID{deposits[n-len(cards)].DepositID}
	}
	return senderFilter, nil
}

func chooseReceiver() (models2.BalanceFilter, error) {
	var ch string
	var receiverFilter models2.BalanceFilter
	fmt.Println("Choose receiver type: ")
	fmt.Println("1.Card\n2.Deposit\n3.Balance id")
	fmt.Scanln(&ch)
	clearConsole()
	fmt.Println("Enter receiver id")
	var e string
	fmt.Scanln(&e)
	id, err := strconv.Atoi(e)
	if err != nil {
		return models2.BalanceFilter{}, err
	}
	switch ch {
	case "1":
		receiverFilter.CardIDs = []models2.CardID{models2.CardID(id)}
	case "2":
		receiverFilter.DepositIDs = []models2.DepositID{models2.DepositID(id)}
	case "3":
		receiverFilter.IDs = []models2.BalanceID{models2.BalanceID(id)}
	}
	return receiverFilter, nil
}
