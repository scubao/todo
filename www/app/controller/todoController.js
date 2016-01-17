myTodoList.controller('todoController', ['$scope', 'entryFactory', function ($scope, entryFactory) {

    $scope.todos;


    getAllTodos();

    function getAllTodos() {
        entryFactory.getEntries()
            .then(function (response) {
                    $scope.todos = response.data;
                },
                function (response) {
                    console.log("Error: " + response.status);
                }
            );

    }

    $scope.addTodo = function(name) {
        entryFactory.createEntry(name)
            .then(function(response) {
                $scope.todoName = '';
                    getAllTodos();
                },
                function (response) {
                    console.log("Error: " + response.status);
                }
            );
    };

    $scope.deleteTodo = function(id) {
        entryFactory.deleteEntry(id)
            .then(function(response) {
                    getAllTodos();
                },
                function (response) {
                    console.log("Error: " + response.status);
                }
            );
    };

    $scope.updateTodo = function(todo) {
        entryFactory.updateEntry(todo)
            .then(function () {
                getAllTodos();
            },
            function (error) {
            $scope.status = 'Unable to update data: ' + error.message;
        });

    }

}]);