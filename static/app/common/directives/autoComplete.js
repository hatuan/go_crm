/**
 * Created by tuanha-01 on 2016-11-09.
 */
define(['angularAMD'], function (angularAMD) {

    var injectParams = ['$http', '$interpolate', '$parse'];

    var autoCompleteDirective = function ($http, $interpolate, $parse) {
        return {
            restrict: 'A',
            require: 'ngModel',
            compile: function (elem, attrs) {
                var
                    modelAccessor = $parse(attrs.ngModel),
                    labelExpression = "{{Code}} - {{Description}}",
                    url = "/api/autocomplete";

                if (attrs.label) {
                    labelExpression = attrs.label;
                }
                if (attrs.url) {
                    url = attrs.url;
                }
                return function (scope, element, attrs) {
                    var
                        mappedItems = null,
                        allowCustomEntry = attrs.allowCustomEntry || false;

                    element.autocomplete({
                        source: function (request, response) {
                            $http({
                                url: url,
                                method: 'GET',
                                params: { object:attrs.object, term: request.term }
                            })
                                .success(function (data) {
                                    mappedItems = $.map(data, function (item) {
                                        var result = {};
                                        if (typeof item === 'string') {
                                            result.label = item;
                                            result.value = item;
                                            return result;
                                        }

                                        result.label = $interpolate(labelExpression)(item);

                                        if (attrs.value) {
                                            result.value = item[attrs.value];
                                        }
                                        else {
                                            result.value = item;
                                        }

                                        return result;
                                    });

                                    return response(mappedItems);
                                });
                        },

                        select: function (event, ui) {
                            scope.$apply(function (scope) {
                                modelAccessor.assign(scope, ui.item.value);
                            });

                            if (attrs.onSelect) {
                                scope.$apply(attrs.onSelect);
                            }

                            element.val(ui.item.label);

                            event.preventDefault();
                        },

                        change: function () {
                            var
                                currentValue = element.val(),
                                matchingItem = null;

                            if (allowCustomEntry) {
                                return;
                            }

                            if (mappedItems) {
                                for (var i = 0; i < mappedItems.length; i++) {
                                    if (mappedItems[i].label === currentValue) {
                                        matchingItem = mappedItems[i].label;
                                        break;
                                    }
                                }
                            }

                            if (!matchingItem) {
                                scope.$apply(function (scope) {
                                    modelAccessor.assign(scope, null);
                                });
                            }
                        }
                    });
                }
            },
        }
    };

    autoCompleteDirective.$inject = injectParams;

    angularAMD.directive('autoComplete', autoCompleteDirective)
});
