// Function to fetch and display employees
function fetchAndDisplayEmployees() {
    fetch('/employees')
        .then(response => response.json())
        .then(data => {
            showEmployeeTable(data);
        })
        .catch(error => console.error('Error fetching employees:', error));
}

// Function to show employee details in a table
function showEmployeeTable(data) {
    // Display employee details in a table
    const table = document.createElement('table');
    table.className = 'table table-bordered';
    table.innerHTML = `
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Email</th>
            </tr>
        </thead>
        <tbody>
            ${data.map(employee => `
                <tr>
                    <td>${employee.ID}</td>
                    <td>${employee.Name}</td>
                    <td>${employee.Email}</td>
                </tr>
            `).join('')}
        </tbody>
    `;

    // Show table in a modal
    const modal = createModal('Employee Details', table);
    document.body.appendChild(modal);
}

// Function to create a modal
function createModal(title, content) {
    const modal = document.createElement('div');
    modal.className = 'modal';
    modal.innerHTML = `
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">${title}</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                </div>
                <div class="modal-body">
                    ${content.outerHTML}
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                </div>
            </div>
        </div>
    `;
    return modal;
}

// Function to fetch employee by ID
function fetchEmployeeById() {
    const employeeId = document.getElementById('employeeId').value;
    fetch(`/employees/${employeeId}`)
        .then(response => response.json())
        .then(data => {
            showEmployeeTable([data]);
            document.getElementById('viewEmployeeForm').style.display = 'none';
        })
        .catch(error => {
            alert(`Error fetching employee by ID: ${error}`);
            console.error('Error fetching employee by ID:', error);
        });

    return false; // Prevent form submission
}

// Function to show the form for adding a new employee
function showAddEmployeeForm() {
    const form = document.createElement('form');
    form.innerHTML = `
        <div class="form-group">
            <label for="newName">Name:</label>
            <input type="text" class="form-control" id="newName" required>
        </div>
        <div class="form-group">
            <label for="newEmail">Email:</label>
            <input type="email" class="form-control" id="newEmail" required>
        </div>
        <button type="button" class="btn btn-primary btn-sm" onclick="addEmployee()">Add Employee</button>
    `;

    // Show form in a modal
    const modal = createModal('Add New Employee', form);
    document.body.appendChild(modal);
}

// Function to add a new employee
function addEmployee() {
    const newName = document.getElementById('newName').value;
    const newEmail = document.getElementById('newEmail').value;

    fetch('/employees', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ Name: newName, Email: newEmail }),
    })
        .then(response => response.json())
        .then(data => {
            console.log('Employee added successfully:', data);
            fetchAndDisplayEmployees(); // Refresh the employee list
        })
        .catch(error => console.error('Error adding employee:', error));

    // Close the modal
    document.querySelector('.modal').remove();
}

// Fetch employees when the page loads
document.addEventListener('DOMContentLoaded', fetchAndDisplayEmployees);

// ... (your existing functions)

// Function to show the form for updating an employee
function showUpdateEmployeeForm() {
    document.getElementById('updateEmployeeForm').style.display = 'block';
}

// Function to update an employee
function updateEmployee() {
    const updateEmployeeId = document.getElementById('updateEmployeeId').value;
    const updateName = document.getElementById('updateName').value;
    const updateEmail = document.getElementById('updateEmail').value;

    fetch(`/employees/${updateEmployeeId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ Name: updateName, Email: updateEmail }),
    })
        .then(response => response.json())
        .then(data => {
            console.log('Employee updated successfully:', data);
            fetchAndDisplayEmployees(); // Refresh the employee list
        })
        .catch(error => console.error('Error updating employee:', error));

    // Close the modal
    document.querySelector('.modal').remove();
}

// Function to show the form for deleting an employee
function showDeleteEmployeeForm() {
    document.getElementById('deleteEmployeeForm').style.display = 'block';
}

// Function to delete an employee
function deleteEmployee() {
    const deleteEmployeeId = document.getElementById('deleteEmployeeId').value;

    fetch(`/employees/${deleteEmployeeId}`, {
        method: 'DELETE',
    })
        .then(response => response.json())
        .then(data => {
            console.log('Employee deleted successfully:', data);
            fetchAndDisplayEmployees(); // Refresh the employee list
        })
        .catch(error => console.error('Error deleting employee:', error));

    // Close the modal
    document.querySelector('.modal').remove();
}
