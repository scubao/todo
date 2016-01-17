/**
 * Created by oliver on 16.01.16.
 */
angular.module('myTodoList', []);

angular.module('myTodoList').factory('todoFactory', ['$http', function ($http) {

    console.log("todoFactory used.")
    var urlBase = "/todo";
    var todoFactory = {};

    todoFactory.getAllTodos = function () {
        return $http.get(urlBase);
    }

    return todoFactory;

}]);

angular.module('myTodoList').controller('todoController',['$scope','entryFactory', function ($scope, todoFactory) {
    $scope.todos = todos;
    $scope.status = "Hi";
}]);

angular.module('myTodoList').controller('MyCtrl', ['$scope', function ($scope) {
    console.log ("Hans");
}]);
