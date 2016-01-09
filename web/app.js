var app = angular.module('todoApp', ['ngMaterial', 'ngResource', 'ngMdIcons']);

app.factory('TodoService', function ($resource) {
    return $resource('http://127.0.0.1:8080/todo/:id', { id: '@id'  }, {
        update: {
            method: 'PUT' // a change request
        }
    });
});

app.controller('TodoCtrl', function ($scope, TodoService) {
    $scope.todos = TodoService.query();

    $scope.AddTodo = function () {
        TodoService.save({
            name: $scope.todoname
        });
        $scope.message = $scope.todoname;
        $scope.todoname = ''
        $scope.todos = TodoService.query();
    };

    $scope.DeleteTodo = function (id) {
        TodoService.delete({
            id
        });
        $scope.todos = TodoService.query();
    };

    $scope.UpdateTodo = function (todo) {
        console.log(todo);
/*
        TodoService.update({
            todo
        });
        $scope.todos = TodoService.query();
*/
    };

});