/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'ajaxService', 'alertsService', 'myApp.Search', 'businessRelationSectorsService'], function (angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', 'alertsService', 'businessRelationSectorsService'];

    var businessRelationSectorsController = function ($scope, $rootScope, $state, $window, moment, alertsService, businessRelationSectorsService) {

        $scope.initializeController = function () {
            $rootScope.applicationModule = "BusinessRelationSectors";
            $rootScope.alerts = [];

            $scope.Search = "";
            $scope.isSearched = false;
            $scope.SortExpression = "Code";
            $scope.SortDirection = "ASC";
            $scope.FetchSize = 100;
            $scope.CurrentPage = 1;
            $scope.PageSize = 9;
            $scope.TotalRows = 0;
            $scope.Selection=[];

            $scope.searchConditionObjects = [];
            $scope.searchConditionObjects.push({
                ID: "business_relation_sector.code",
                Name: "Code",
                Type: "CODE", //CODE, FREE, DATE
                ValueIn: "BusinessRelationSector",
                Value: ""
            },
            {
                ID: "business_relation_sector.name",
                Name: "Name",
                Type: "FREE", //CODE, FREE, DATE
                ValueIn: "",
                Value: ""
            });

            $scope.BusinessRelationSectors = [];
            $scope.FilteredItems = [];
            $scope.getBusinessRelationSectors();
        };

        $scope.refresh = function () {
            $scope.getBusinessRelationSectors();
        }

        $scope.showSearch = function () {
            $scope.isSearched = !$scope.isSearched;
        }

        $scope.selectAll = function () {
            $scope.Selection=[];
            for(var i = 0; i < $scope.FilteredItems.length; i++) {
                $scope.Selection.push($scope.FilteredItems[i]["ID"]);
            }
        }

        $scope.delete = function () {
            if($scope.Selection.length <= 0)
                return;
            var deleteBusinessRelationSectors = $scope.createDeleteBusinessRelationSectorObject()
            businessRelationSectorsService.deleteBusinessRelationSector(deleteBusinessRelationSectors, 
                function (response, status) {
                    $scope.getBusinessRelationSectors();
                }, 
                function (response, status){
                    alertsService.RenderErrorMessage(response.Error);
                });    
        }

        $scope.toggleSelection = function (_id) {
             var idx = $scope.Selection.indexOf(_id);
             if (idx > -1) {
               $scope.Selection.splice(idx, 1);
             }
             else {
               $scope.Selection.push(_id);
             }
        };

        $scope.getBusinessRelationSectors = function (searchSqlCondition) {
            if(!angular.isUndefinedOrNull(searchSqlCondition))
                $scope.Search = searchSqlCondition;
            var businessRelationSectorInquiry = $scope.createBusinessRelationSectorObject();
            businessRelationSectorsService.getBusinessRelationSectors(businessRelationSectorInquiry, $scope.businessRelationSectorsInquiryCompleted, $scope.businessRelationSectorsInquiryError);
        };

        $scope.businessRelationSectorsInquiryCompleted = function (response, status) {
            alertsService.RenderSuccessMessage(response.ReturnMessage);
            $scope.BusinessRelationSectors = response.Data.BusinessRelationSectors;
            $scope.TotalRows = response.TotalRows;
            $scope.Selection = [];
            $scope.FilteredItems = [];
        };

        $scope.businessRelationSectorsInquiryError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.createBusinessRelationSectorObject = function () {
            var businessRelationSectorInquiry = new Object();

            businessRelationSectorInquiry.Search = $scope.Search;
            businessRelationSectorInquiry.SortExpression = $scope.SortExpression;
            businessRelationSectorInquiry.SortDirection = $scope.SortDirection;
            businessRelationSectorInquiry.FetchSize = $scope.FetchSize;

            return businessRelationSectorInquiry;
        }

        $scope.createDeleteBusinessRelationSectorObject = function() {
            var deleteBusinessRelationSectors = new Object();
            deleteBusinessRelationSectors.ID = $scope.Selection.join(",");
            return deleteBusinessRelationSectors;
        }
    };

    businessRelationSectorsController.$inject = injectParams;
    angularAMD.controller('BusinessRelationSectorsController', businessRelationSectorsController);
});
