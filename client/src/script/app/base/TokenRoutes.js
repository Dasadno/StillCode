document.addEventListener('DOMContentLoaded', async () => {
const token = localStorage.getItem('token');
  const FeaturesButton = document.getElementById('FeaturesButton');
  const ProblemsButton = document.getElementById('ProblemsButton');
  const CommunityButton = document.getElementById('CommunityButton');

  if (!token) {
    authButtons.style.display = 'flex';
    userWidget.style.display  = 'none';

    FeaturesButton.href = '/signin';
    ProblemsButton.href = '/signin';
    CommunityButton.href = '/signin';
  
    return;
  }

});