angular.module('todoApp', ['ngMaterial', 'restangular', 'ngMdIcons']);

angular.module('todoApp').config(function (RestangularProvider) {
    RestangularProvider.setRestangularFields({
        id: "_id"
    });
});

// Inject Restangular into your controller
angular.module('todoApp').controller('TodoCtrl', function ($scope, Restangular) {

    //    var Todos = Restangular.all('http://127.0.0.1:8080/todo');
    var Todos = Restangular.all('todo');
    var allTodos = Todos.getList();
    allTodos.then(function (data) {
        $scope.todos = data;
    })

    $scope.AddTodo = function () {};

    $scope.DeleteTodo = function () {};

    $scope.UpdateTodo = function () {};

});