/**
 * Created by tuanha-01 on 5/6/2016.
 */
"use strict";

define(['angularAMD', 'jquery', 'jquery.validate', 'jquery.validation.extend', 'bootstrap', 'ui-bootstrap', 'kendo.all.min', 'kendo.culture.en', 'kendo.culture.us', 'kendo.culture.vi', 'kendo.culture.vn', 'angular-validate', 'angular-globalize-wrapper', 'jquery-validation-globalize', 'ui.router', 'satellizer', 'pascalprecht.translate', 'blockUI', 'stateConfig', 'toastr', 'angular-moment', 'ngInfiniteScroll', 'bootstrap-switch', 'angular-bootstrap-switch', 'angular-confirm-modal', 'myApp.navBar', 'myApp.Capitalize', 'myApp.Constants'], function(angularAMD) {

    var app = angular.module("myApp", ['ui.router', 'satellizer', 'pascalprecht.translate', 'blockUI', 'toastr', 'angularMoment', 'ui.bootstrap', 'kendo.directives', 'ngValidate', 'globalizeWrapper', 'infinite-scroll', 'frapontillo.bootstrap-switch', 'angular-confirm', 'myApp.NavBar', 'myApp.Capitalize', 'myApp.Constants']);

    app.config(function(blockUIConfig) {

        // Change the default overlay message
        blockUIConfig.message = "executing...";
        // Change the default delay to 100ms before the blocking is visible
        blockUIConfig.delay = 1;
        // Disable automatically blocking of the user interface
        blockUIConfig.autoBlock = false;

    });

    app.config(['$authProvider', function($authProvider) {
        // Satellizer configuration that specifies which API
        // route the JWT should be retrieved from
        $authProvider.loginUrl = '/api/token-auth';
        $authProvider.loginRoute = '/login';
        $authProvider.tokenName = "Token";
    }]);

    //https://github.com/Foxandxss/angular-toastr
    app.config(function(toastrConfig) {
        angular.extend(toastrConfig, {
            allowHtml: true,
            closeButton: true,
            closeHtml: '<button>&times;</button>',
            extendedTimeOut: 1000,
            timeOut: 5000,
        });
    });

    app.config(stateConfig);

    app.config(['$validatorProvider', function($validatorProvider) {
        $validatorProvider.setDefaults({
            highlight: function(element) {
                $(element).closest('.form-group').addClass('has-error');
            },
            unhighlight: function(element) {
                $(element).closest('.form-group').removeClass('has-error');
            },
            errorElement: 'span',
            errorClass: 'has-block',
            errorPlacement: function(error, element) {
                return true;

                //if (element.parent('.input-group').length) {
                //    error.insertAfter(element.parent());
                //} else {
                //    error.insertAfter(element);
                //}
            },
            /* http://stackoverflow.com/questions/21813868/adding-jquery-validation-to-kendo-ui-elements */
            /* http://lukaszledochowski.blogspot.nl/2015/02/validation-using-aspnet-mvc-kendo-ui.html */
            /* https://github.com/jpkleemans/angular-validate */
            ignore: []
        });
    }]);


    app.config(['globalizeWrapperProvider', function(globalizeWrapperProvider) {
        // The path to cldr-data
        globalizeWrapperProvider.setCldrBasePath('bower_components/cldr-data');

        // The path to messages
        globalizeWrapperProvider.setL10nBasePath('l10n');

        // Files to load in main dir: "{{cldrBasePath}}/main/{{locale}}"
        globalizeWrapperProvider.setMainResources([
            'currencies.json',
            'ca-gregorian.json',
            'timeZoneNames.json',
            'numbers.json'
        ]);

        // Files to load in supplemental dir: "{{cldrBasePath}}/supplemental'
        globalizeWrapperProvider.setSupplementalResources([
            'currencyData.json',
            'likelySubtags.json',
            'plurals.json',
            'timeData.json',
            'weekData.json'
        ]);
    }]);

    app.controller('indexController', ['$scope', '$rootScope', '$http', 'blockUI', function($scope, $rootScope, $http, blockUI) {

        $scope.initializeController = function() {
            $rootScope.displayContent = false;
            // if ($location.path() != "") {
            $scope.initializeApplication($scope.initializeApplicationComplete, $scope.initializeApplicationError);
            // }
        };

        $scope.initializeApplicationComplete = function(response) {
            $rootScope.MenuItems = response.MenuItems;
            $rootScope.displayContent = true;
            $rootScope.IsInitAppCompleted = true;
        };

        $scope.initializeApplicationError = function(response) {
            alert("ERROR : InitializeApplication");
        };

        $scope.initializeApplication = function(successFunction, errorFunction) {
            blockUI.start();
            $scope.AjaxGet("/api/main/initializeApplication", successFunction, errorFunction);
            blockUI.stop();
        };

        $scope.AjaxGet = function(route, successFunction, errorFunction) {
            setTimeout(function() {
                $http({ method: 'GET', url: route }).success(function(response, status, headers, config) {
                    successFunction(response, status);
                }).error(function(response) {
                    errorFunction(response);
                });
            }, 1);

        };

        $scope.AjaxGetWithData = function(data, route, successFunction, errorFunction) {
            setTimeout(function() {
                $http({ method: 'GET', url: route, params: data }).success(function(response, status, headers, config) {
                    successFunction(response, status);
                }).error(function(response) {
                    errorFunction(response);
                });
            }, 1);

        }

    }]);

    app.run(['$state', '$rootScope', '$auth', 'globalizeWrapper', 'amMoment', function($state, $rootScope, $auth, globalizeWrapper, amMoment) {

        //kendo.culture("vi-VN");

        // kendo-date-picker config
        //$rootScope.datePickerConfig = {
        //   format: "dd/MM/yyyy",
        //   parseFormats: ["yyyy-MM-dd", "dd/MM/yyyy", "yyyy/MM/dd"],
        //};

        $rootScope.isAuthenticated = function() {
            return $auth.isAuthenticated();
        };

        globalizeWrapper.loadLocales(['vi', 'en']);

        $rootScope.$on('GlobalizeLoadSuccess', function() {
            //console.log("GlobalizeLoadSuccess"); 
        });

        $rootScope.$on('GlobalizeLocaleChanged', function() {
            //console.log("globalizeWrapper.getLocale() = " + globalizeWrapper.getLocale());
            Globalize.locale(globalizeWrapper.getLocale());
        });
    }]);

    angular.isUndefinedOrNull = function(val) {
        return angular.isUndefined(val) || val === null;
    }

    // Bootstrap Angular when DOM is ready
    angularAMD.bootstrap(app);

    return app;
});