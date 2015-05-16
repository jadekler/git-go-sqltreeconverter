$(document).ready(function() {
    $("#addCondition input[type='submit']").click(function(e) {
        e.preventDefault();
        addCondition();
        updateConditions();
        setupDraggables();
    });

    setupResetConditions();
    setupUpdateConditions();
    setupDraggables();
});

function setupResetConditions() {
    $(".resetConditions").click(function() {
        resetConditions($(this).attr('data-reset-type'));
    });
}

function setupUpdateConditions() {
    $("#updateConditions").click(updateConditions);
}

function setupDraggables() {
    var sortable = $("#sortable");
    sortable.sortable();
    sortable.disableSelection();
}

function resetConditions(resetType) {
    $.ajax({
        url: "/reset",
        method: "PUT",
        cache: false,
        data: {
            resetType: resetType
        },
        success: function(data) {
            data = JSON.parse(data);

            redrawData(data['tree']);
            redrawUsers(data['allUsers']);
            rematchUsers(data['matchingUsers']);
        }
    });
}

function redrawUsers(allUsers) {
    var tableBody = $(".user-sql-rows tbody");
    tableBody.empty();

    allUsers.forEach(function(user) {
        $(".user-sql-rows tbody").append("<tr id="+user.Id+"><td>"+user.Name+"</td><td>"+user.Age+"</td><td>"+user.NumPets+"</td></tr>");
    });
}

function rematchUsers(matchingUsers) {
    $(".user-sql-rows tbody tr").each(function(key, elem) {
        elemIsMatchingUser = $.grep(matchingUsers, function(e){ return e.Id == $(elem).attr('id'); }).length > 0;
        
        if (elemIsMatchingUser) {
            $(elem).css('background-color', 'green');
        } else {
            $(elem).css('background-color', 'red');
        }
    });
}

function redrawData(treeData) {
    $.ajax({
        url: "/",
        method: "GET",
        cache: false,
        success: function(data) {
            var pageData = $.parseHTML(data);
            $("#sortable").replaceWith($(pageData).find("#sortable"));
            $(".condition-sql-rows").replaceWith($(pageData).find(".condition-sql-rows"));
            setupDraggables();
        }
    });

    initTree(treeData);
}

function updateConditions() {
    var conditions = [];
    $("#sortable li").each(function(k, v) {
        var condition = {
            Type: $(v).attr("data-type"),
            Text: $(v).text(),
            Field: $(v).attr("data-field"),
            Operator: $(v).attr("data-operator"),
            Value: $(v).attr("data-value")
        }

        conditions.push(condition);
    });

    $.ajax({
        url: "/conditions",
        method: "PUT",
        data: {
            conditions: JSON.stringify(conditions)
        },
        success: function(data) {
            data = JSON.parse(data);

            redrawData(data['tree']);
            rematchUsers(data['matchingUsers']);
        }
    });
}

function addCondition() {
    var value = $("#value");

    if (!value.val()) {
        return
    }

    var field = $("#field").val();
    var operator = $("#operator").val();

    var logicTemplate = $("#sortable .logic").first().clone();
    logicTemplate.text("AND");

    var equalityTemplate = $("#sortable .equality").first().clone();
    equalityTemplate.attr("data-field", field);
    equalityTemplate.attr("data-operator", operator);
    equalityTemplate.attr("data-value", value);
    equalityTemplate.text(field + " " + operator + " " + value);

    $("#sortable .scope").last().before(logicTemplate).before(equalityTemplate);

    value.val("");
}