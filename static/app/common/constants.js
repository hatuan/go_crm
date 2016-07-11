/**
 * Created by tuanha-01 on 5/11/2016.
 */
"use strict";

define(['angular'], function (angular) {
    var constantModule = angular.module('myApp.Constants', [])
    constantModule.constant('Constants', {
        Status: [
            {Code : 0, Name: 'InActive'},
            {Code : 1, Name: 'Active'},
        ]
    })
});