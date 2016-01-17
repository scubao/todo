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

angular.module('myTodoList').controller("todoController", ['$scope', 'entryFactory', function ($scope, todoFactory) {

        $scope.todos;

        getAllTodos();

        function getAllTodos() {
            todoFactory.getAllTodos()
                .success(function (todos) {
                    $scope.todos = todos;
                })
                .error(function (error) {
                    $scope.status = 'Unable to load customer data: ' + error.message;
                })

        };

}]);