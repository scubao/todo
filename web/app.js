var app = angular.module('todoApp', ['ngMaterial', 'ngResource', 'ngMdIcons']);
//var app = angular.module('todoApp', ['ngResource']);

app.factory('TodoService', function ($resource) {
    return $resource('http://127.0.0.1:8080/todo/:id', {}, {
        'update': {
            method: 'PUT',
            params: {
                id: '@id',
                todo: '@todo'
            }
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
        id = todo.id;
        entry = todo;
        TodoService.update({
            id
        }, {
            entry
        });
        console.log(id);
        console.log(entry);
        $scope.todos = TodoService.query();
    };

});