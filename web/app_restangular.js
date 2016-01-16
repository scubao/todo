angular.module('todoApp', []);

angular.module('todoApp').factory('todoFactory', ['$http', function ($http) {
    var urlBase = '/todo';
    var todoFactory = {};

    todoFactory.getTodos = function () {
        return $http.get(urlBase);
    };

    todoFactory.deleteTodo = function (id) {
        return $http.delete(urlBase + '/' + id);
    }

    todoFactory.updateTodo = function (todo) {
        return $http.put(urlBase + '/' + todo.id, todo)
    }

    return todoFactory;
}]);

angular.module('todoApp').controller('todoController', ['$scope', 'todoFactory', function ($scope, todoFactory) {

    $scope.todos;
    $scope.status;

    var todoItem = {};

    getTodos();
    updateTodo(todoItem);

    function getTodos() {
        todoFactory.getTodos()
            .success(function (todos) {
                $scope.todos = todos;
            })
            .error(function (error) {
                $scope.status = 'Error loading Data: ' + error.message;
            })
    };

    function updateTodo(todoItem) {
        todoFactory.updateTodo(todoItem)
            .success(function () {
                $scope.status = 'todoItem Updated';
            })
            .error(function (error) {
                $scope.status = 'Error ' + error.message;
            })
    };

}]);