/**
 * Created by tuanha-01 on 5/30/2016.
 */
"use strict";

define(['angular'], function (angular) {

    angular.module('myApp.moment', []).factory('moment', ['$window', function ($window) {
        var _moment = function () {
            return $window.moment;
        };
        return {
            moment: _moment
        }
    }]);

});