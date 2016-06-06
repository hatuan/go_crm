/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['application-configuration'], function (app) {

    var injectParams = ['$scope', '$rootScope'];

    var MasterController = function ($scope, $rootScope) {
        $scope.initializeController = function () {
            $scope.activeTab = 1;
        };
        $scope.setActiveTab = function(tabToSet) {
            $scope.activeTab = tabToSet;
        };
    }

    MasterController.$inject = injectParams;

    app.register.controller('MasterController', MasterController);
});
