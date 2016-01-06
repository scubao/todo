var app = angular.module('todoApp', ['ngMaterial', 'ngResource']);
app.factory('Entry', function ($resource) {
    return $resource('http://127.0.0.1:8080/todo/:id');
});

app.controller('ResourceController', function ($scope, Entry) {
    var blabla = Entry.query();
    $scope.blabla = blabla;
    console.log(blabla);
});


app.controller('todoCtrl', function ($scope) {
    $scope.todos = [
        {
            "id": "5686419788d8e71954cb3ef3",
            "created": 1451639191877627021,
            "name": "Gasflasche",
            "done": true
            }, {
            "id": "5686419788d8e71954cb3ef4",
            "created": 1451639191877627145,
            "name": "Bierkasten",
            "done": false
            }
    ];
});