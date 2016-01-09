var app = angular.module('todoApp', ['ngMaterial', 'ngResource', 'ngMdIcons']);

app.factory('TodoService', function ($resource) {
    return $resource('http://127.0.0.1:8080/todo/:id');
});

app.controller('TodoCtrl', function ($scope, TodoService) {
    $scope.todos = TodoService.query();

    $scope.ButtonClick = function () {
        TodoService.save({ name: $scope.todoname});
        $scope.message = $scope.todoname;
        $scope.todoname = ''
        $scope.todos = TodoService.query();
    };
    
});
