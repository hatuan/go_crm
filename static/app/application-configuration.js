/**
 * Created by tuanha-01 on 5/6/2016.
 */
"use strict";

define(['angularAMD', 'ui.router', 'satellizer', 'pascalprecht.translate', 'blockUI', 'stateConfig', 'jquery', 'bootstrap', 'toastr' ], function (angularAMD) {
    var app = angular.module("myApp", ['ui.router', 'satellizer', 'pascalprecht.translate', 'blockUI', 'toastr']);

    app.config(function (blockUIConfig) {

        // Change the default overlay message
        blockUIConfig.message = "executing...";
        // Change the default delay to 100ms before the blocking is visible
        blockUIConfig.delay = 1;
        // Disable automatically blocking of the user interface
        blockUIConfig.autoBlock =false;

    });

    app.config(['$authProvider', function($authProvider) {
        // Satellizer configuration that specifies which API
        // route the JWT should be retrieved from
        $authProvider.loginUrl = '/api/token-auth';
        $authProvider.loginRoute = '/login';
    }]);

    app.config(stateConfig);

    app.controller('indexController', ['$scope', '$rootScope', '$http', 'blockUI', function ($scope, $rootScope, $http, blockUI) {

        $scope.initializeController = function () {
            $rootScope.displayContent = false;
            // if ($location.path() != "") {
                $scope.initializeApplication($scope.initializeApplicationComplete, $scope.initializeApplicationError);
            // }
        };

        $scope.initializeApplicationComplete = function (response) {
            $rootScope.MenuItems = response.MenuItems;
            $rootScope.displayContent = true;
            $rootScope.IsInitAppCompleted = true;
        };

        $scope.initializeApplicationError = function (response) {
            alert("ERROR : InitializeApplication");
        };

        $scope.initializeApplication = function (successFunction, errorFunction) {
            blockUI.start();
            $scope.AjaxGet("/api/main/initializeApplication", successFunction, errorFunction);
            blockUI.stop();
        };

        $scope.AjaxGet = function (route, successFunction, errorFunction) {
            setTimeout(function () {
                $http({method: 'GET', url: route}).success(function (response, status, headers, config) {
                    successFunction(response, status);
                }).error(function (response) {
                    errorFunction(response);
                });
            }, 1);

        };

        $scope.AjaxGetWithData = function (data, route, successFunction, errorFunction) {
            setTimeout(function () {
                $http({method: 'GET', url: route, params: data}).success(function (response, status, headers, config) {
                    successFunction(response, status);
                }).error(function (response) {
                    errorFunction(response);
                });
            }, 1);

        }

    }]);

    // Bootstrap Angular when DOM is ready
    angularAMD.bootstrap(app);

    return app;
});


