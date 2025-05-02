package handlers

import (
	"context"
	"task_manager/internal/taskService"
	"task_manager/internal/userService"
	"task_manager/internal/web/tasks"

	"github.com/google/uuid"
)

type TaskHandler struct {
	taskService taskService.TaskService
	userService userService.UserService
}

func NewTaskHandler(ts taskService.TaskService, us userService.UserService) *TaskHandler {
	return &TaskHandler{
		taskService: ts,
		userService: us,
	}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.taskService.GetAllTasks()
	if err != nil {
		return nil, err
	}

	respone := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     (*uuid.UUID)(&tsk.ID),
			Task:   &tsk.Task,
			IsDone: &tsk.Is_done,
		}
		respone = append(respone, task)
	}

	return respone, nil
}

func (h *TaskHandler) GetTasksByUserID(_ context.Context, request tasks.GetTasksByUserIDRequestObject) (tasks.GetTasksByUserIDResponseObject, error) {
	userTasks, err := h.userService.GetTasksForUser(request.Id)
	if err != nil {
		return nil, err
	}

	respone := tasks.GetTasksByUserID200JSONResponse{}

	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:     (*uuid.UUID)(&tsk.ID),
			Task:   &tsk.Task,
			IsDone: &tsk.Is_done,
			UserId: &tsk.UserID,
		}
		respone = append(respone, task)
	}

	return respone, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:    *taskRequest.Task,
		Is_done: *taskRequest.IsDone,
	}
	createdTask, err := h.taskService.CreateTask(taskToCreate, *taskRequest.UserId)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.Is_done,
	}

	return response, nil
}

func (h *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskRequest := request.Body
	taskID := request.Id

	taskToUpdate := taskService.Task{
		Task:    *taskRequest.Task,
		Is_done: *taskRequest.IsDone,
	}
	updatedTask, err := h.taskService.UpdateTask(taskID, taskToUpdate)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.Is_done,
	}

	return response, nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	if err := h.taskService.DeleteTask(taskID); err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}
