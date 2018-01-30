package apis

import (
	"encoding/json"
	"io/ioutil"
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
	ID *datastore.Key `datastore:"__key__" json:"id,omitempty"`

	// From where to inherit
	// Required: true
	ParentID string `datastore:"-" json:"parent_id,omitempty"`

	// First Name of the User
	// Required: true
	Firstname string `datastore:"firstname,omitempty" json:"firstname,omitempty"`

	// Middle Name of the User
	Middlename string `datastore:"middlename,omitempty" json:"middlename,omitempty"`

	// Last Name of the User
	Lastname string `datastore:"lastname,omitempty" json:"lastname,omitempty"`

	// Validated Email Ids
	EmailIds []string `datastore:"email_ids,omitempty" json:"email_ids,omitempty"`

	// 10 Digit Unformatted Phone Numbers
	PhoneNumbers []string `datastore:"phone_numbers,omitempty" json:"phone_numbers,omitempty"`

	// Date of Birth in UTC Zulu Format
	Dob time.Time `datastore:"dob,omitempty" json:"dob,omitempty"`

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

const (
	ISO8601 = "2006-01-02T15:04:05.000Z"
)

func (u *User) MarshalJSON() ([]byte, error) {
	type marshalledUser User
	return json.Marshal(&struct {
		Dob         string `datastore:"dob" json:"dob"`
		DateCreated string `datastore:"date_created" json:"date_created"`
		DateUpdated string `datastore:"date_updated" json:"date_updated"`
		*marshalledUser
	}{
		DateCreated:    time.Time(u.DateCreated).UTC().Format(ISO8601),
		DateUpdated:    time.Time(u.DateUpdated).UTC().Format(ISO8601),
		Dob:            time.Time(u.Dob).UTC().Format(ISO8601),
		marshalledUser: (*marshalledUser)(u),
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, "Failed to establish connection with Database: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	b, _ := ioutil.ReadAll(r.Body)

	var user = new(User)

	err = json.Unmarshal(b, user)
	if err != nil {
		http.Error(w, "Unable to parse Request BODY: "+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	parentKey, err := datastore.DecodeKey(user.ParentID)
	if err != nil {
		http.Error(w, "Unable to Decode Parent ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	key := datastore.IDKey("Users", 0, parentKey)
	key.Namespace = "NeverLand"

	user.DateCreated = time.Now().UTC()
	user.DateUpdated = time.Now().UTC()

	key, err = client.Put(ctx, key, user)
	if err != nil {
		http.Error(w, "Failed to create User: "+err.Error(), http.StatusExpectationFailed)
		return
	}

	user.ID = key

	json, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Unable to convert response to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	params := mux.Vars(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, "Failed to establish connection with Database: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	key, err := datastore.DecodeKey(params["id"])
	if err != nil {
		http.Error(w, "Unable to Decode ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	key.Namespace = "NeverLand"

	if err = client.Delete(ctx, key); err != nil {
		http.Error(w, "Failed to Delete User: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, "Failed to establish connection with Database: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	query := datastore.NewQuery("Users").Limit(10).Namespace("NeverLand")

	var users []*User

	keys, err := client.GetAll(ctx, query, &users)
	if err != nil {
		http.Error(w, "No Users Found: "+err.Error(), http.StatusNotFound)
		return
	}

	json, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Unable to convert response to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(keys) > 0 {
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	params := mux.Vars(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, "Failed to establish connection with Database: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	var user = new(User)

	key, err := datastore.DecodeKey(params["id"])
	if err != nil {
		http.Error(w, "Unable to Decode ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = client.Get(ctx, key, user)
	if err != nil {
		http.Error(w, "User Not Found: "+err.Error(), http.StatusNotFound)
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Unable to convert response to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	params := mux.Vars(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, "Failed to establish connection with Database: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	key, err := datastore.DecodeKey(params["id"])
	if err != nil {
		http.Error(w, "Unable to Decode ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	key.Namespace = "NeverLand"

	b, _ := ioutil.ReadAll(r.Body)

	var newUser = new(User)

	err = json.Unmarshal(b, newUser)
	if err != nil {
		http.Error(w, "Unable to parse Request BODY: "+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var oldUser User

	_, err = client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		if err := tx.Get(key, &oldUser); err != nil {
			return err
		}

		dateCreated, _ := time.Parse(ISO8601, time.Time(oldUser.DateCreated).UTC().Format(ISO8601))
		newUser.DateCreated = dateCreated
		newUser.DateUpdated = time.Now().UTC()

		oldUser = *newUser
		_, err := tx.Put(key, &oldUser)
		return err
	})
	if err != nil {
		http.Error(w, "Failed to Update: "+err.Error(), http.StatusExpectationFailed)
		return
	}

	json, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, "Unable to convert response to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UploadProfileImage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
