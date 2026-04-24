# 🌴 Bienvenidos a la Jungla de Bits 🐒💾

> *"En esta selva, los árboles son de silicio, las lianas son de fibra óptica y los monos programan en sus ratos libres."*

¡Saludos, explorador digital! Has llegado a **Hola Jungla de Bits**, el rincón más salvaje del ecosistema binario, donde los `0` y los `1` conviven en perfecta (y caótica) armonía. Ajusta tu sombrero de safari, carga tu mochila con buenos commits y prepárate para una expedición épica.

## Endpoint Go: informacion del host

Este repo ahora incluye una pequena API HTTP en Go que expone informacion del host y tambien puede usarse en modo CLI.

### Ejecutar como servidor

```bash
go run .
```

El servidor levanta por defecto en `:8080` y expone:

- `GET /api/host-info` -> respuesta JSON para webapps.
- `GET /api/host-info?format=text` -> respuesta texto plano para CLI.
- `GET /health` -> estado simple del servicio.

Ejemplos:

```bash
curl http://localhost:8080/api/host-info
curl http://localhost:8080/api/host-info?format=text
```

### Ejecutar como CLI

```bash
go run . -cli
go run . -cli -format=text
```

La salida entrega:

- hostname
- sistema operativo
- arquitectura
- cantidad de CPU
- version de Go
- IPs locales detectadas
- fecha/hora de recoleccion en UTC

---

## 🗺️ ¿Qué es esta jungla?

Este repositorio es el punto de partida de una aventura por el vasto follaje del código. Aquí no hay mapas perfectos, pero sí muchas pistas, ramas (`git branch`) colgando por todos lados y algún que otro `bug` escondido entre las hojas esperando ser cazado.

- 🌿 **Exploración libre:** cada archivo es un nuevo sendero.
- 🐍 **Cuidado con las serpientes:** también conocidas como *null pointers*.
- 🦜 **Los loros repiten logs:** préstales atención, a veces dicen la verdad.
- 🐘 **La memoria es grande:** pero no infinita, ¡libera lo que uses!

---

## 🧭 Primeros pasos en la selva

1. **Clona tu machete** (digo, el repo):
   ```bash
   git clone https://github.com/adminevolsystemcl/hola-jungla-de-bits.git
   ```
2. **Entra en la espesura:**
   ```bash
   cd hola-jungla-de-bits
   ```
3. **Respira hondo.** El aire huele a café recién hecho y a compilador caliente. ☕

---

## 🐾 Fauna local

| Criatura            | Hábitat                     | Peligro |
|---------------------|-----------------------------|----------|
| 🐒 Mono Refactor    | Ramas muy largas            | Medio   |
| 🐛 Bug Escurridizo  | Cualquier línea no testeada | Alto    |
| 🦋 Feature Bonita   | Issues abiertos             | Bajo    |
| 🐊 Deadline         | Viernes por la tarde        | Extremo |

---

## 🌋 Reglas de supervivencia

- Nunca hagas `push --force` sin avisar a la tribu. 🗣️
- Si ves un `TODO` de hace tres años, déjalo descansar… o adóptalo con cariño. 🪴
- Los *merge conflicts* se resuelven con paciencia, no con machete.
- Un buen `commit message` vale más que mil stickers.

---

## 🤝 Únete a la expedición

¿Quieres dejar tu huella en el barro digital? ¡Perfecto!

- Abre un **issue** si encuentras una criatura extraña.
- Manda un **pull request** si domesticaste alguna.
- Comparte la aventura con otros exploradores.

---

## 🔥 Palabras finales alrededor de la fogata

La **Jungla de Bits** no es solo código: es curiosidad, comunidad y un poquito de locura binaria. Así que respira, teclea y disfruta. Recuerda:

> *"No todo el que deambula por la jungla está perdido… a veces solo está debuggeando."* 🧉✨

¡Feliz hacking y que los bits te acompañen! 🦍💚
