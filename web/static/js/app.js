document.addEventListener("DOMContentLoaded", () => {
    fetch("/api/artists")
      .then(res => res.json())
      .then(artists => {
        let container = document.getElementById("artists");
        artists.forEach(artist => {
          let card = document.createElement("div");
          card.className = "card";
          card.innerHTML = `
            <img src="${artist.image}" alt="${artist.name}">
            <h3>${artist.name}</h3>
            <p>🎤 Año inicio: ${artist.year}</p>
            <p>💿 Primer álbum: ${artist.firstAlbum}</p>
            <p>👥 Miembros: ${artist.members.join(", ")}</p>
          `;
          container.appendChild(card);
        });
      })
      .catch(err => console.error("Error cargando artistas:", err));
  });
  