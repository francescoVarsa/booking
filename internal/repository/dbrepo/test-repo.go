package dbrepo

import (
	"booking/internal/models"
	"errors"
	"time"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InserReservation insert a reservation in the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// if the room id is 2, then fail; otherwise, pass
	if res.RoomID == 2 {
		return 0, errors.New("some error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("Some error")
	}
	return nil
}

// SearchAvailabilityByDates returns true if availability exists for room id and false if no availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

//SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil

}

// GetRoomById gets a room by id
func (m *testDBRepo) GetRoomById(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("Some error")
	}

	return room, nil

}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 1, "", nil
}

func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

// AllNewReservations returns a slice of all reservations
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

func (m *testDBRepo) GetReservationById(id int) (models.Reservation, error) {
	var res models.Reservation

	return res, nil
}

func (m *testDBRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

func (m *testDBRepo) DeleteReservation(id int) error {

	return nil
}

func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {

	return nil
}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}
