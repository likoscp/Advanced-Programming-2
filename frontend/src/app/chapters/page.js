
"use client";

import { useEffect, useState, Suspense } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";

function ChaptersContent() {
    const [chapters, setChapters] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true);

    const router = useRouter();

    useEffect(() => {
        const fetchChapters = async () => {
            setLoading(true);
            const token = localStorage.getItem("token");

            if (!token) {
                setError("Token not found, please sign in.");
                window.location.href = '/sign-in';
                setLoading(false);
                return;
            }

            try {
                const response = await axios.get("http://localhost:8089/chapters/", {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                if (Array.isArray(response.data)) {
                    setChapters(response.data);
                } else {
                    setError("Invalid response structure.");
                }


            } catch (error) {
                if (error.response && error.response.status === 401) {
                    window.location.href = '/sign-in';
                } else {
                    setError("Error loading chapters: " + error.message);
                }
            }
            setLoading(false);
        };

        fetchChapters();
    }, []);

    return (
        <div>
            <h1>Chapter List</h1>
            {error && <p>{error}</p>}
            {loading ? (
                <p>Loading...</p>
            ) : chapters.length === 0 ? (
                <p>No data.</p>
            ) : (
                <ul>
                    {chapters.map((chapter) => (
                        <li key={chapter.id}>
                            <br />
                            ID: {chapter.id} <br />
                            Title: {chapter.title} <br />
                            Number: {chapter.number} <br />
                            Created At: {new Date(chapter.created_at).toLocaleDateString()} <br />
                        </li>
                    ))}

                </ul>
            )}
        </div>
    );
}

export default function Chapters() {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            <ChaptersContent />
        </Suspense>
    );
}
