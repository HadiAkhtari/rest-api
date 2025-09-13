package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"
)

func (app *application) createEvent(ctx *gin.Context) {
	var event database.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := app.models.Events.Insert(&event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, event)
}

func (app *application) getAllEvents(ctx *gin.Context) {
	events, err := app.models.Events.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to get all events"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}
func (app *application) getEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to get event"})
		return
	}
	event, err := app.models.Events.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to get event"})
		return
	}
	if event == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	ctx.JSON(http.StatusOK, event)

}
func (app *application) updateEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to get event"})
		return
	}
	existingEvent, err := app.models.Events.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to get event"})
		return
	}
	if existingEvent == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	updateEvent := &database.Event{}
	if err := ctx.ShouldBindJSON(&updateEvent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Faild to update event"})
		return
	}
	updateEvent.Id = id
	if err := app.models.Events.Update(updateEvent); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to update event"})
		return
	}
	ctx.JSON(http.StatusOK, updateEvent)

}
func (app *application) deleteEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Faild to get event"})
		return
	}
	if err := app.models.Events.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to delete event"})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Event deleted"})
}

func (app application) addAttendeeToEvent(ctx *gin.Context) {
	eventIdStr := ctx.Param("id")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	userIdStr := ctx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	event, err := app.models.Events.Get(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to get event"})
		return
	}
	if event == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	userToAdd, err := app.models.Users.Get(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to get user to add"})
		return
	}
	if userToAdd == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	existingAttendee, err := app.models.Attendees.GetByEventAndAttendee(event.Id, userToAdd.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to get user to add"})
		return
	}
	if existingAttendee != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Attendee already exists"})
		return
	}
	attendee := database.Attendee{
		Event_Id: eventId,
		User_Id:  userId,
	}
	_, err = app.models.Attendees.Insert(&attendee)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to add attendee"})
		return
	}
	ctx.JSON(http.StatusCreated, attendee)
}

func (app *application) getAttendeesForEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	users, err := app.models.Attendees.GetAttendeesByEvent(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)

}
func (app *application) deleteAttendeeFromEvent(ctx *gin.Context) {
	eventIdStr := ctx.Param("id")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	userIdStr := ctx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = app.models.Attendees.Delete(userId, eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Attendee deleted"})
}

func (app *application) getEventsByAttendee(ctx *gin.Context) {
	IdStr := ctx.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	events, err := app.models.Attendees.GetByAttendee(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)

}
