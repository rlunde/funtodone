/*global Backbone */
'use strict';

// FunToDone is global for developing in the console and functional testing.
var App = Backbone.Marionette.Application.extend({
});

window.FunToDone = new App();

(function () {
	var filterState = new Backbone.Model({
		filter: 'all'
	});

	FunToDone.reqres.setHandler('filterState', function () {
		return filterState;
	});
})();

FunToDone.on('before:start', function () {
	// FunToDone.setRootLayout();
});
