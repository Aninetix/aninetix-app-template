# Aninetix App Template

🚀 **Aninetix App Template** est un template de base pour créer rapidement des applications modulaires basées sur le framework **Aninetix Core** (https://github.com/Aninetix/core)

Ce template fournit une architecture prête à l’emploi pour développer des applications composées de **modules indépendants**, communiquant entre eux via un **EventBus middleware** supportant :

* 🔄 **asynchrone (Send)**
* 🔁 **synchrone (SendSync – request/reply)**

Chaque module a un **objectif précis**, isolé, testable et remplaçable.

---

## 🧠 Concept

L’application est construite autour de :

* **AnWare (core runtime)**
* **Modules** chargés automatiquement
* **EventBus central** pour la communication inter-modules
* **Middleware interne** pour le dispatch, le routage et le lifecycle

### Communication inter-modules

* 🔹 `Send` → envoi non bloquant
* 🔹 `SendSync` → envoi bloquant avec réponse (style TCP interne)

---

## 📦 Contenu du template

### 🔹 Module `anConsol` (inclus par défaut)

Un module console déjà initialisé permettant :

* d’interagir avec l’application via des commandes
* d’envoyer des événements aux autres modules
* de tester rapidement les flux inter-modules

👉 Ce module est **spécifique à l’app finale** et peut être modifié ou supprimé selon le besoin.

---

### 🔹 Dossier `anparam` (obligatoire)

Le dossier `anparam` contient les structures partagées nécessaires au démarrage de l’application :

* configuration du Core
* flags globaux
* paramètres communs aux modules

⚠️ Ce dossier est requis pour toute application basée sur Aninetix.

---

## 🗂 Structure du projet

```
.
├── _cmd/
│   └── init/            # Script d'initialisation du template
├── anmodules/           # Modules Aninetix
│   └── anConsol/        # Module console par défaut
├── anparam/             # Paramètres & configuration (obligatoire)
├── go.mod
├── go.sum
└── README.md
```
---

## ⚙️ Initialisation du projet

Après avoir cloné le template, il est nécessaire de **l’initialiser** afin de :

* renommer le module Go
* mettre à jour tous les imports
* préparer le projet pour une nouvelle application

### 1️⃣ Cloner le template

```
git clone [https://github.com/Aninetix/aninetix-app-template](https://github.com/Aninetix/aninetix-app-template) my-app
cd my-app
```

---

### 2️⃣ Initialiser le projet

```
go run ./_cmd/init <nouveau-nom-du-module>
```

Exemple :

```
go run ./_cmd/init github.com/Aninetix/my-awesome-app
```

Cela va automatiquement :

* ✅ mettre à jour `go.mod`
* ✅ réécrire tous les imports Go
* ✅ préparer le projet pour un nouveau repo

---

### 3️⃣ Finaliser

```
rm -rf .git
git init
go mod tidy
```

---

## 🧩 Développement des modules

Chaque module :

* implémente l’interface `AnModule`
* possède son propre cycle de vie
* communique via l’EventBus Aninetix
* peut être **synchrone ou asynchrone**

Les modules peuvent être :

* internes à l’application
* ou distribués via des repositories séparés (`go get`)

---

## 🎯 Objectif

Ce template a pour but de :

* accélérer la création d’applications Aninetix
* imposer une architecture claire et scalable
* favoriser la modularité et l’interopérabilité
* permettre une évolution indépendante des modules

---

## 🧬 Basé sur

* **Aninetix Core**
  Framework modulaire basé sur un middleware EventBus

---

## 🛠 Licence

MIT

---

💡 Pour toute contribution ou amélioration, n’hésitez pas à ouvrir une issue ou une pull request.
