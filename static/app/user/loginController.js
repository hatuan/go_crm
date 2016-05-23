/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['application-configuration', 'alertsService'], function (app) {
    var injectParams = ['$scope', '$rootScope', '$auth', 'alertsService', 'toastr', '$state', '$http', '$window'];

    var LoginController = function ($scope, $rootScope, $auth, alertsService, toastr, $state, $http, $window) {

        $rootScope.closeAlert = alertsService.closeAlert;
        $rootScope.alerts = [];

        $scope.initializeController = function () {

            $scope.UserName = "";
            $scope.Password = "";

            alertsService.RenderSuccessMessage("Please register if you do not have an account.");

        };

        $scope.login = function () {
            $rootScope.IsloggedIn = false;
            var credentials = $scope.createLoginCredentials();
            $auth.login(credentials)
                .then(function () {
                    // Return an $http request for the now authenticated
                    // user so that we can flatten the promise chain
                    return $http.get("/api/token-auth");
                    //Handle error
                }, function (error) {
                    // Because we returned the $http.get request in the $auth.login
                    // promise, we can chain the next promise to the end here
                })
                .then(function (response) {
                    toastr.success('You have successfully signed in!');

                    $window.localStorage.setItem('currentUser', JSON.stringify(response.data));

                    setTimeout(function () {
                        $state.go('home');
                    }, 10);
                })
                .catch(function (error) {
                    $scope.clearValidationErrors();
                    toastr.error(error.data.message, error.status);
                });
        };

        $scope.clearValidationErrors = function () {

            $scope.UserNameInputError = false;
            $scope.PasswordInputError = false;

        };

        $scope.createLoginCredentials = function () {

            var user = new Object();

            user.username = $scope.UserName;
            user.password = $scope.Password;

            return user;

        }

    };

    LoginController.$inject = injectParams;
    app.register.controller('LoginController', LoginController);
});
