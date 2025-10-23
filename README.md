# 🌐 Portfolio – Valentine Ladjyn

Mon portfolio personnel développé en **Go**, **HTML**, **CSS** et **JavaScript**, conçu pour présenter mes projets et compétences en développement web full-stack.  

🔗 **En ligne :** [valentineladjyn.eruditaoue.fr](https://valentineladjyn.eruditaoue.fr)

---

## 🧩 Sections principales
- **Navigation :** barre fixe avec accès rapide aux différentes sections et téléchargement du CV  
- **Hero :** introduction animée avec mon nom et rôle  
- **À propos :** courte présentation de mon parcours  
- **Compétences :** visualisation circulaire avec pourcentages et technologies  
- **Projets :** cartes interactives avec effet *flip* et carousel horizontal  
- **Contact :** formulaire dynamique relié à ma messagerie et lien vers mon profil LinkedIn  
- **Crédits :** mention de l’auteur (© 2025 Valentine Ladjyn)

---

## ⚙️ Stack technique

**Langages principaux :**
- Go (serveur)
- HTML / CSS / JavaScript (frontend)

**Package externe :**
- [`github.com/joho/godotenv`](https://github.com/joho/godotenv) – gestion des variables d’environnement

**Architecture :**
- Backend en Go servant les pages statiques et gérant le formulaire de contact (via SMTP)
- Frontend responsive, animé en CSS et JS vanilla
- Hébergement sur **VPS personnel** (Nginx + PM2)

---

## 🚀 Installation et exécution

### 1. Cloner le projet
```bash
git clone https://github.com/<ton-utilisateur>/portfolio.git
cd portfolio
```

### 2. Créer un fichier .env
```bash
SMTP_HOST=smtp.exemple.com
SMTP_PORT=587
SMTP_USER=ton.email@exemple.com
SMTP_PASS=ton_mot_de_passe
```

### 3. Lancer le serveur
```bash
go run . + port (optionnel)
```

Le site est ensuite accessible sur http://localhost:5280 (ou port choisi)


## 🧠 Compétences mises en avant

- Développement **full-stack** (Go / HTML / CSS / JS)
- Gestion d’un serveur HTTP en Go
- Manipulation d’environnements et variables sécurisées
- Intégration d’un **formulaire de contact fonctionnel**
- Création d’une **interface dynamique et responsive**
- Déploiement sur **serveur VPS (Nginx + PM2)**

## 👩‍💻 Auteur

**Valentine Ladjyn**
            
[🔗 LinkedIn](https://www.linkedin.com/in/valentine-ladjyn-182020329?lipi=urn%3Ali%3Apage%3Ad_flagship3_messaging_conversation_detail%3BkUto6tCEQkGBOeQ0546Xsw%3D%3D)

[📧 Me contacter via le site](https://valentineladjyn.eruditaoue.fr/#contact)

© 2025 Valentine Ladjyn – Tous droits réservés.