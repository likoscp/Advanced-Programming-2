
"use client";

import { useEffect, useState, Suspense } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";

function ComicsContent() {
    const [comics, setComics] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true);

    const router = useRouter();

    useEffect(() => {
        const fetchComics = async () => {
            setLoading(true);
            const token = localStorage.getItem("token");

            if (!token) {
                setError("Token not found, please sign in.");
                window.location.href = '/sign-in';
                setLoading(false);
                return;
            }

            try {
                const response = await axios.get("http://localhost:8089/comics/", {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                if (Array.isArray(response.data)) {
                    setComics(response.data);
                } else {
                    setError("Invalid response structure.");
                }


            } catch (error) {
                if (error.response && error.response.status === 401) {
                    window.location.href = '/sign-in';
                } else {
                    setError("Error loading comics: " + error.message);
                }
            }
            setLoading(false);
        };

        fetchComics();
    }, []);

    return (
        <div>
            <h1>Comic List</h1>
            {error && <p>{error}</p>}
            {loading ? (
                <p>Loading...</p>
            ) : comics.length === 0 ? (
                <p>No data.</p>
            ) : (
                <ul>
                    {comics.map((comic) => (
                        <li key={comic.id}>
                            <br />
                            ID: {comic.id} <br />
                            Title: {comic.title} <br />
                            Description: {comic.description} <br />
                        </li>
                    ))}

                </ul>
            )}
        </div>
    );
}

export default function Comics() {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            <ComicsContent />
        </Suspense>
    );
}
