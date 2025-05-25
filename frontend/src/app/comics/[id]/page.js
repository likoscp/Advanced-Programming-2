"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";

export default function ComicPage() {
  const params = useParams(); 
  const [comic, setComic] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!params.id) return;

    const fetchComic = async () => {
      try {
        const response = await fetch(`http://localhost:8089/comics/${params.id}`);
        if (!response.ok) {
          throw new Error("Failed to fetch comic");
        }
        const data = await response.json();
        setComic(data);
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };

    fetchComic();
  }, [params.id]);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;
  if (!comic) return <p>Comic not found</p>;

  return (
    <div>
      <h1>Comic Details</h1>
      <ul>
        <li key={comic._id}>
          ID: {comic._id} <br />
          Title: {comic.title} <br />
          Author ID: {comic.author_id} <br />
          Translator ID: {comic.translator_id} <br />
          Artist ID: {comic.artist_id} <br />
          Description: {comic.description} <br />
          Cover Image: <img src={comic.cover_image} alt={comic.title} style={{ maxWidth: "200px" }} /> <br />
          Status: {comic.status} <br />
          Release Date: {new Date(comic.comic_date).toLocaleDateString()} <br />
          Created At: {new Date(comic.created_at).toLocaleDateString()} <br />
          Updated At: {new Date(comic.updated_at).toLocaleDateString()} <br />
          Views: {comic.views} <br />
          Rating: {comic.rating} <br />
          Chapters: {comic.chapters && comic.chapters.length > 0 ? comic.chapters.map(chapter => chapter.title).join(", ") : "None"} <br />

        </li>
      </ul>
    </div>
  );
}
