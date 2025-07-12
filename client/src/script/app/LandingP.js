 document.addEventListener('DOMContentLoaded', () => {
    const icon = document.getElementById('profileIcon');
    const buttons = document.getElementById('authButtons');

    if (localStorage.getItem('token')) {
        document.getElementById('authButtons').style.display = 'none';
        document.getElementById('profileIcon').style.display = 'block';
    }
});


