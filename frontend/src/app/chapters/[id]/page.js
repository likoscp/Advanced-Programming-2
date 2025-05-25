"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";

export default function ChapterPage() {
  const params = useParams(); 
  const [chapter, setChapter] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!params.id) return;

    const fetchChapter = async () => {
      try {
        const response = await fetch(`http://localhost:8089/chapters/${params.id}`);
        if (!response.ok) {
          throw new Error("Failed to fetch chapter");
        }
        const data = await response.json();
        setChapter(data);
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };

    fetchChapter();
  }, [params.id]);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;
  if (!chapter) return <p>Chapter not found</p>;

  return (
    <div>
      <h1>Chapter Details</h1>
      <ul>
        <li key={chapter._id}>
          ID: {chapter._id} <br />
          Comic ID: {chapter.comic_id} <br />
          Title: {chapter.title} <br />
          Number: {chapter.number} <br />
          Created At: {new Date(chapter.created_at).toLocaleDateString()} <br />
          Likes: {chapter.likes} <br />
          Dislikes: {chapter.dislikes} <br />
          Pages: {chapter.pages && chapter.pages.length > 0 ? chapter.pages.map(page => (
            <div key={page.id}>
              Page {page.page_num}: <img src={page.image_url} alt={`Page ${page.page_num}`} style={{ maxWidth: "200px" }} />
            </div>
          )) : "None"} <br />
          <br />


        </li>
      </ul>
    </div>
  );
}

