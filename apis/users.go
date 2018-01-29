package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

var projectID = "chrome-setup-158308"

type User struct {
	// Auto Generated
	// Read Only: true
	ID *datastore.Key `datastore:"__key__" json:"id"`

	// From where to inherit
	// Required: true
	ParentID string `datastore:"-" json:"parent_id"`

	// First Name of the User
	// Required: true
	Firstname string `datastore:"firstname" json:"firstname"`

	// Middle Name of the User
	Middlename string `datastore:"middlename,omitempty" json:"middlename,omitempty"`

	// Last Name of the User
	Lastname string `datastore:"lastname,omitempty" json:"lastname,omitempty"`

	// Validated Email Ids
	EmailIds []string `datastore:"email_ids" json:"email_ids"`

	// 10 Digit Unformatted Phone Numbers
	PhoneNumbers []string `datastore:"phone_numbers" json:"phone_numbers"`

	// Date of Birth in UTC Zulu Format
	// Required: true
	Dob time.Time `datastore:"dob" json:"dob"`

	// Role of the User
	Role string `datastore:"role,omitempty" json:"role,omitempty"`

	// User's Profile Photo, Needs to be Uploaded
	// Read Only: true
	AvatarURL string `datastore:"avatar_url,omitempty" json:"avatar_url,omitempty"`

	// Blood Group of the User
	BloodGroup string `datastore:"blood_group,omitempty" json:"blood_group,omitempty"`

	// Date in UTC Zulu Format
	// Read Only: true
	DateCreated time.Time `datastore:"date_created" json:"date_created"`

	// Date in UTC Zulu Format
	// Read Only: true
	DateUpdated time.Time `datastore:"date_updated" json:"date_updated"`

	// Status of the User's Account
	// Read Only: true
	Status string `datastore:"status,omitempty" json:"status,omitempty"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type marshalledUser User
	return json.Marshal(&struct {
		Dob string `datastore:"dob" json:"dob"`
		*marshalledUser
	}{
		Dob:            time.Time(u.Dob).Format("2006-01-02"),
		marshalledUser: (*marshalledUser)(u),
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		fmt.Printf("Failed")
	}

	b, _ := ioutil.ReadAll(r.Body)

	var user = new(User)

	err = json.Unmarshal(b, user)
	if err != nil {
		log.Fatalf("Failed to Unmarshall: %v", err)
		fmt.Printf("Failed")
	}

	grandParentKey := datastore.IDKey("Society", 5066549580791808, nil)
	grandParentKey.Namespace = "NeverLand"
	parentKey := datastore.IDKey("Family", 5668600916475904, grandParentKey)
	parentKey.Namespace = "NeverLand"
	// parentKey, err := datastore.DecodeKey(*user.ParentID)
	// if err != nil {
	// 	w.WriteHeader(http.StatusFailedDependency)
	// 	log.Fatalf("Unable to Decode Key: %v", err)
	// }

	key := datastore.IDKey("Users", 0, parentKey)
	key.Namespace = "NeverLand"

	key, err = client.Put(ctx, key, user)
	if err != nil {
		log.Fatalf("Failed to Create: %v", err)
		fmt.Printf("Failed")
	}

	user.ParentID = key.String()

	json, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write(json)
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		fmt.Printf("Failed")
	}

	query := datastore.NewQuery("Users").Limit(10).Namespace("NeverLand")

	var users []*User

	keys, err := client.GetAll(ctx, query, &users)
	if err != nil {
		w.Write([]byte("Query Failed: " + err.Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if len(keys) > 0 {
		w.Write(json)
		w.WriteHeader(http.StatusOK)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	params := mux.Vars(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		fmt.Printf("Failed")
	}

	var user = new(User)

	key, err := datastore.DecodeKey(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		log.Fatalf("Unable to Decode Key: %v", err)
	}

	err = client.Get(ctx, key, user)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		log.Fatalf("Unable to Get: %v", err)
	}

	json, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Bad Result: %v", err)
		w.WriteHeader(http.StatusGone)
	}
	w.Write(json)
	w.WriteHeader(http.StatusOK)
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UploadProfileImage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
