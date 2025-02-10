package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type Profile struct {
	Mobile                         string `json:"mobile"`
	Name                           string `json:"name"`
	StudentID                      string `json:"studentId"`
	StatusCode                     int    `json:"statusCode"`
	CommunicationCoordinatorStatus *int   `json:"communicationCoordinatorStatus"`
	TechnicalCoordinatorStatus     int    `json:"technicalCoordinatorStatus"`
	OnHold                         bool   `json:"onHold"`
	UnderQa                        bool   `json:"underQa"`
	Week                           int    `json:"week"`
}

type Student struct {
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
	StudentID  string `json:"studentId"`
	StatusCode int    `json:"statusCode"`
}

type BrototypeUser struct {
	Profile Profile `json:"profile"`
	Student Student `json:"student"`
	Token   string  `json:"token"`
	Msg     string  `json:"msg"`
}
type Domain struct {
	Typename string `json:"__typename"`
	Name     string `json:"name"`
	ID       string `json:"id"`
}

type ProfileStudent struct {
	ID                     string  `json:"_id"`
	TenthStatus            bool    `json:"10th_status"`
	TwelfthStatus          bool    `json:"12th_status"`
	AadharNumber           string  `json:"adaar_nmbr"`
	Address                string  `json:"address"`
	AdmissionCounselor     string  `json:"admission_counsellor"`
	Approved               bool    `json:"approved"`
	Batch                  string  `json:"batch"`
	Branch                 string  `json:"branch"`
	District               string  `json:"district"`
	DOB                    string  `json:"dob"`
	Domain                 Domain  `json:"domain"`
	Email                  string  `json:"email"`
	Father                 string  `json:"father"`
	Fumigation             bool    `json:"fumigation"`
	Gender                 string  `json:"gender"`
	Guardian               string  `json:"guardian"`
	GuardianAddress        string  `json:"guardian_address"`
	GuardianContact        string  `json:"guardian_contact"`
	GuardianRelation       string  `json:"guardian_rel"`
	HigherEduStatus        bool    `json:"higherEdu_status"`
	Hub                    string  `json:"hub"`
	Mobile                 string  `json:"mobile"`
	Mother                 string  `json:"mother"`
	Name                   string  `json:"name"`
	Pincode                string  `json:"pincode"`
	Place                  string  `json:"place"`
	State                  string  `json:"state"`
	StudentID              string  `json:"studentId"`
	Update                 bool    `json:"update"`
	ZohoContactID          string  `json:"zoho_contact_id"`
	AdmissionID            string  `json:"admissionId"`
	StudentDetailsID       string  `json:"studentDetailsId"`
	Background             *string `json:"background"`
	EducationQualification *string `json:"educationQualification"`
	EducationStatus        *string `json:"educationStatus"`
	EducationSubject       *string `json:"educationSubject"`
	FatherContact          *string `json:"fatherContact"`
	MotherContact          *string `json:"motherContact"`
	WorkExperience         *string `json:"workExperience"`
}

type Response struct {
	Msg     string         `json:"msg"`
	Status  bool           `json:"status"`
	Student ProfileStudent `json:"student"`
}
type Review struct {
	ID              string     `json:"id"`
	Week            string     `json:"week"`
	PreferredTime   time.Time  `json:"preferredTime"`
	ReviewStageCode int        `json:"reviewStageCode"`
	Ratings         []Rating   `json:"ratings"`
	ScheduledOn     *time.Time `json:"scheduledOn"`
	Status          *string    `json:"status"`
	CompletedOn     *time.Time `json:"completedOn"`
	Advisor         string     `json:"advisor"`
	MeetLink        *string    `json:"meetLink"`
	ReviewType      string     `json:"reviewType"`
	ReviewBadge     *string    `json:"reviewBadge"`
	SpecialType     *string    `json:"specialType"`
}

type Rating struct {
	ReviewID         string `json:"reviewId"`
	ID               string `json:"_id"`
	Rating           int    `json:"rating"`
	ReviewedEntityID string `json:"reviewedEntityId"`
	EntityType       int    `json:"entityType"`
	Comment          string `json:"comment"`
	ReviewType       int    `json:"reviewType"`
}

type ReviewsResponse struct {
	Reviews []Review `json:"reviews"`
}

func authenticateWithBrototype(url string, user User) (*BrototypeUser, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("error occured")
		return nil, err
	}
	resp, err := http.Post(url+"/auth/signIn", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("error occured in http request")
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != fiber.StatusOK {
		return nil, fmt.Errorf("failed to make request")

	}
	var brototypeUser BrototypeUser
	if err := json.NewDecoder(resp.Body).Decode(&brototypeUser); err != nil {
		fmt.Printf("failed to decode json")
	}
	return &brototypeUser, nil
}

var url string = "https://api.brototype.com/tool/student/api"

func GetUserAuth(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	brototypeUser, err := authenticateWithBrototype(url, user)
	if err != nil {
		fmt.Printf("failed to send http request")
	}
	//so what happens here is pretty self explanator the GetUser post request returns the brototype user
	return c.JSON(brototypeUser)
}

func GetUserDetails(c *fiber.Ctx) error {
	tokenWithBearer := c.Get("Authorization")
	if tokenWithBearer == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "token is empty"})
	}
	token := strings.Split(tokenWithBearer, " ")[1]
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "token is empty"})
	}
	req, err := http.NewRequest("GET", url+"/profile/details", nil)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to create request"})
	}
	req.Header.Set("Authorization", tokenWithBearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to send req"})
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != fiber.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	var profileStudentResponse Response
	if err := json.NewDecoder(resp.Body).Decode(&profileStudentResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to convert json"})
	}
	return c.Status(fiber.StatusOK).JSON(profileStudentResponse)
}

func BrototypeReviews(url string, token string) (*ReviewsResponse, error) {
	req, err := http.NewRequest("GET", url+"/review?avoidReviewType=4&pageSize=100", nil)
	if err != nil {
		fmt.Println("failed to create request")
		return nil, err
	}
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("failed to send request")
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != fiber.StatusOK {
		return nil, fmt.Errorf("failed to get reviews")
	}
	var reviewsReponse ReviewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&reviewsReponse); err != nil {
		return nil, fmt.Errorf("failed to decode json")
	}
	return &reviewsReponse, nil

}

func GetUserReviews(c *fiber.Ctx) error {
	tokenWithBearer := c.Get("Authorization")
	if tokenWithBearer == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "token is empty"})
	}
	fmt.Println(tokenWithBearer)
	reviewResponse, err := BrototypeReviews(url, tokenWithBearer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get reviews"})
	}
	return c.Status(fiber.StatusOK).JSON(reviewResponse)
}
