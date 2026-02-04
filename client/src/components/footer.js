/**
 * Footer Component - Reusable site footer using DOM methods
 */

/**
 * Create footer element using DOM methods
 * @param {Object} options - Footer options
 * @param {boolean} options.minimal - Minimal footer (just copyright)
 */
export function createFooter(options = {}) {
  const { minimal = false } = options;

  const footer = document.createElement('footer');
  footer.className = minimal ? 'bg-darker py-4' : 'bg-dark py-12';

  const container = document.createElement('div');
  container.className = 'container mx-auto px-4';

  if (minimal) {
    container.className += ' text-center text-gray-500 text-sm';
    const p = document.createElement('p');
    p.textContent = '\u00A9 2025 StillCode. All rights reserved.';
    container.appendChild(p);
    footer.appendChild(container);
    return footer;
  }

  // Main grid container
  const grid = document.createElement('div');
  grid.className = 'grid grid-cols-1 md:grid-cols-4 gap-8 mb-8';

  // Brand section
  const brandSection = createBrandSection();
  grid.appendChild(brandSection);

  // Resources section
  const resourcesSection = createLinkSection('Resources', [
    { href: '/problems', text: 'Problems' },
    { href: '#', text: 'Contests' },
    { href: '#', text: 'Practice' },
    { href: '#', text: 'Discussion' },
    { href: '#', text: 'Articles' }
  ]);
  grid.appendChild(resourcesSection);

  // About section
  const aboutSection = createLinkSection('About', [
    { href: '#', text: 'About Us' },
    { href: '#', text: 'Our Team' },
    { href: '#', text: 'Press' },
    { href: '#', text: 'Contact' },
    { href: '#', text: 'Blog' }
  ]);
  grid.appendChild(aboutSection);

  // Support section
  const supportSection = createLinkSection('Support', [
    { href: '#', text: 'Help Center' },
    { href: '#', text: 'FAQs' },
    { href: '#', text: 'Privacy Policy' },
    { href: '#', text: 'Terms of Service' },
    { href: '#', text: 'Cookie Policy' }
  ]);
  grid.appendChild(supportSection);

  container.appendChild(grid);

  // Copyright section
  const copyright = document.createElement('div');
  copyright.className = 'border-t border-gray-800 pt-8 text-center text-gray-500 text-sm';
  const copyrightP = document.createElement('p');
  copyrightP.textContent = '\u00A9 2025 StillCode. All rights reserved.';
  copyright.appendChild(copyrightP);
  container.appendChild(copyright);

  footer.appendChild(container);
  return footer;
}

/**
 * Create the brand section with logo and social links
 */
function createBrandSection() {
  const section = document.createElement('div');

  // Logo
  const logoContainer = document.createElement('div');
  logoContainer.className = 'flex items-center mb-4';

  const logoIcon = document.createElement('div');
  logoIcon.className = 'text-primary text-2xl font-bold mr-2';
  const icon = document.createElement('i');
  icon.className = 'fas fa-code';
  logoIcon.appendChild(icon);

  const logoH3 = document.createElement('h3');
  logoH3.className = 'text-xl font-bold';
  logoH3.textContent = 'Still';
  const logoSpan = document.createElement('span');
  logoSpan.className = 'text-primary';
  logoSpan.textContent = 'Code';
  logoH3.appendChild(logoSpan);

  logoContainer.appendChild(logoIcon);
  logoContainer.appendChild(logoH3);
  section.appendChild(logoContainer);

  // Description
  const desc = document.createElement('p');
  desc.className = 'text-gray-400 mb-4';
  desc.textContent = 'Forge your path to coding excellence with our algorithmic challenges.';
  section.appendChild(desc);

  // Social links
  const socialContainer = document.createElement('div');
  socialContainer.className = 'flex space-x-4';

  const socialIcons = ['fa-twitter', 'fa-linkedin', 'fa-github', 'fa-discord'];
  socialIcons.forEach(iconClass => {
    const link = document.createElement('a');
    link.href = '#';
    link.className = 'text-gray-400 hover:text-primary transition-colors';
    const i = document.createElement('i');
    i.className = 'fab ' + iconClass;
    link.appendChild(i);
    socialContainer.appendChild(link);
  });

  section.appendChild(socialContainer);
  return section;
}

/**
 * Create a link section (Resources, About, Support)
 */
function createLinkSection(title, links) {
  const section = document.createElement('div');

  const h4 = document.createElement('h4');
  h4.className = 'text-lg font-semibold mb-4';
  h4.textContent = title;
  section.appendChild(h4);

  const ul = document.createElement('ul');
  ul.className = 'space-y-2';

  links.forEach(link => {
    const li = document.createElement('li');
    const a = document.createElement('a');
    a.href = link.href;
    a.className = 'text-gray-400 hover:text-primary transition-colors';
    a.textContent = link.text;
    li.appendChild(a);
    ul.appendChild(li);
  });

  section.appendChild(ul);
  return section;
}

/**
 * Inject footer into a container element
 */
export function injectFooter(containerId = 'footer-container', options = {}) {
  const container = document.getElementById(containerId);
  if (container) {
    container.appendChild(createFooter(options));
  }
}

export default { createFooter, injectFooter };
