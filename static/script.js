        // Compétences avec pourcentages
        const skills = [
            { name: 'HTML/CSS', percentage: 30 },
            { name: 'Docker', percentage: 30 },
            { name: 'Git', percentage: 75 },
            { name: 'Go', percentage: 60 },
            { name: 'Canva', percentage: 70 }
        ];

        // Projets
        const projects = [
            {
                name: 'Groupie Tracker',
                image: 'img/groupieTracker.png',
                description: 'Application web affichant des artistes et leurs concerts à partir d’une API. Interface dynamique, événements interactifs et backend robuste en Go.',
                skills: ['Go (Serveur Mux, routes HTTP)', 'HTML/CSS', 'JSON', 'APIRest']
            },
            {
                name: 'Petites Victoires',
                image: 'img/Forum.png',
                description: 'Création d’un forum web complet en Go avec authentification, gestion des posts, commentaires, likes et filtres, stockant les données dans une base SQLite.',
                skills: ['Go (session, cookies)', 'HTML/CSS', 'SQLite/SQL', 'Dockerisation et sécurité (bcrypt)']
            },
            {
                name: 'NetCat',
                image: 'img/chris-ried-ieic5Tq8YMk-unsplash.jpg',
                description: 'Développement d’un chat TCP en Go inspiré de NetCat, permettant à plusieurs clients de se connecter à un serveur, d’échanger des messages horodatés et d’être notifiés des connexions et déconnexions en temps réel.',
                skills: ['Go (TCP, goroutines, channels, Mutex)']
            }
        ];

        // ==================== NAVIGATION ==================== 
        
        // Mobile menu toggle
        const mobileToggle = document.getElementById('mobileToggle');
        const navLinks = document.getElementById('navLinks');

        mobileToggle.addEventListener('click', () => {
            mobileToggle.classList.toggle('active');
            navLinks.classList.toggle('active');
        });

        // Active link on scroll
        const sections = document.querySelectorAll('section');
        const navLinksArray = document.querySelectorAll('.nav-link');

        window.addEventListener('scroll', () => {
            let current = '';
            sections.forEach(section => {
                const sectionTop = section.offsetTop;
                const sectionHeight = section.clientHeight;
                if (window.pageYOffset >= sectionTop - 100) {
                    current = section.getAttribute('id');
                }
            });

            navLinksArray.forEach(link => {
                link.classList.remove('active');
                if (link.getAttribute('href').slice(1) === current) {
                    link.classList.add('active');
                }
            });

            // Nav shadow on scroll
            const navbar = document.getElementById('navbar');
            if (window.scrollY > 50) {
                navbar.classList.add('scrolled');
            } else {
                navbar.classList.remove('scrolled');
            }
        });

        // Close mobile menu on link click
        navLinksArray.forEach(link => {
            link.addEventListener('click', () => {
                mobileToggle.classList.remove('active');
                navLinks.classList.remove('active');
            });
        });

        // ==================== COMPÉTENCES CIRCLES ====================
        
        function generateSkills() {
            const skillsGrid = document.getElementById('skillsGrid');
            
            skills.forEach((skill, index) => {
                const radius = 56;
                const circumference = 2 * Math.PI * radius;
                const offset = circumference - (skill.percentage / 100) * circumference;

                const skillItem = document.createElement('div');
                skillItem.className = 'skill-item';
                skillItem.style.animationDelay = `${index * 0.1}s`;
                
                skillItem.innerHTML = `
                    <div class="skill-circle">
                        <svg width="120" height="120">
                            <circle class="bg-circle" cx="60" cy="60" r="${radius}"></circle>
                            <circle class="progress-circle" cx="60" cy="60" r="${radius}"
                                    style="stroke-dasharray: ${circumference}; stroke-dashoffset: ${circumference};"
                                    data-offset="${offset}"></circle>
                        </svg>
                        <div class="skill-percentage">${skill.percentage}%</div>
                    </div>
                    <div class="skill-name">${skill.name}</div>
                `;
                
                skillsGrid.appendChild(skillItem);
            });

            // Animate circles on scroll
            const observer = new IntersectionObserver((entries) => {
                entries.forEach(entry => {
                    if (entry.isIntersecting) {
                        const progressCircles = entry.target.querySelectorAll('.progress-circle');
                        progressCircles.forEach(circle => {
                            const offset = circle.getAttribute('data-offset');
                            circle.style.strokeDashoffset = offset;
                        });
                    }
                });
            }, { threshold: 0.5 });

            observer.observe(skillsGrid);
        }

        // ==================== PROJETS CAROUSEL ====================
        
        function generateProjects() {
            const carousel = document.getElementById('projectsCarousel');
            
            projects.forEach(project => {
                const projectCard = document.createElement('div');
                projectCard.className = 'project-card';
                
                projectCard.innerHTML = `
                    <div class="card-inner">
                        <div class="card-front">
                            <img src="${project.image}" alt="${project.name}">
                            <h3>${project.name}</h3>
                        </div>
                        <div class="card-back">
                            <h3>${project.name}</h3>
                            <p>${project.description}</p>
                            <div class="skills-used">
                                ${project.skills.map(skill => `<span class="skill-tag">${skill}</span>`).join('')}
                            </div>
                        </div>
                    </div>
                `;
                
                carousel.appendChild(projectCard);
            });
        }

        // Carousel navigation
        let currentIndex = 0;
        const prevBtn = document.getElementById('prevBtn');
        const nextBtn = document.getElementById('nextBtn');

        function updateCarousel() {
            const carousel = document.getElementById('projectsCarousel');
            const cardWidth = carousel.querySelector('.project-card').offsetWidth;
            const gap = 30;
            const offset = currentIndex * (cardWidth + gap);
            carousel.style.transform = `translateX(-${offset}px)`;
        }

        prevBtn.addEventListener('click', () => {
            if (currentIndex > 0) {
                currentIndex--;
                updateCarousel();
            }
        });

        nextBtn.addEventListener('click', () => {
            const maxIndex = projects.length - 3;
            if (currentIndex < maxIndex) {
                currentIndex++;
                updateCarousel();
            }
        });

        // Update carousel on window resize
        window.addEventListener('resize', () => {
            currentIndex = 0;
            updateCarousel();
        });

        // ==================== FORMULAIRE CONTACT ====================
        
        const contactForm = document.getElementById('contactForm');
        const formMessage = document.getElementById('formMessage');

        contactForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            
            const formData = {
                name: document.getElementById('name').value,
                email: document.getElementById('email').value,
                message: document.getElementById('message').value
            };

            try {
                const response = await fetch('/api/contact', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(formData)
                });

                if (response.ok) {
                    formMessage.textContent = 'Message envoyé avec succès !';
                    formMessage.className = 'form-message success';
                    contactForm.reset();
                } else {
                    throw new Error('Erreur lors de l\'envoi');
                }
            } catch (error) {
                formMessage.textContent = 'Erreur lors de l\'envoi du message. Veuillez réessayer.';
                formMessage.className = 'form-message error';
            }

            setTimeout(() => {
                formMessage.style.display = 'none';
            }, 5000);
        });

        // ==================== INIT ====================
        
        document.addEventListener('DOMContentLoaded', () => {
            generateSkills();
            generateProjects();
        });




