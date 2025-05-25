"use client";

import React, { useState } from "react";
import Link from "next/link";

const Footer = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  const openModal = () => setIsModalOpen(true);
  const closeModal = () => setIsModalOpen(false);

  return (
    <footer className="bg-gray-800 text-white p-4 text-center">
      <div className="flex flex-col md:flex-row justify-between items-center max-w-4xl mx-auto">
        <a
          href="https://github.com/likoscp/"
          target="_blank"
          rel="noopener noreferrer"
          className="text-gray-400 hover:text-white transition"
        >
          © 2025 SE-2320
        </a>

        <div className="flex gap-4 mt-2 md:mt-0">
          <Link
            href="/docs/privacyPage"
            className="text-gray-400 hover:text-white transition"
          >
            Privacy policy
          </Link>
          <button onClick={openModal} className="text-gray-400 hover:text-white transition">
            Contact us / Report
          </button>
        </div>
      </div>

      {isModalOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-gray-900 p-6 rounded-lg shadow-lg w-96">
            <span className="absolute top-2 right-4 text-gray-600 text-xl cursor-pointer" onClick={closeModal}>
              &times;
            </span>
            <h2 className="text-center text-xl font-bold mb-4">Contact Us</h2>

            <form id="contactForm" className="flex flex-col gap-3">
              <input className="p-2 border rounded" type="text" placeholder="Name..." required />
              <input className="p-2 border rounded" type="email" placeholder="Email..." required />
              <textarea className="p-2 border rounded" placeholder="Message..." required></textarea>
              <button className="bg-blue-600 text-white p-2 rounded hover:bg-blue-700 transition">
                Send
              </button>
            </form>
          </div>
        </div>
      )}
    </footer>
  );
};

export default Footer;
