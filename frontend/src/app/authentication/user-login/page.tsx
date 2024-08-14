"use client";

import React, { useState } from "react";
import { Button, Label } from "flowbite-react";
import Image from "next/image";
import logo from "../../../../public/logo.png";

export default function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const handleLogin = (e) => {
    e.preventDefault();
    setError("");

    // Client-side validation
    if (!username || !password) {
      setError("Both fields are required.");
      return;
    }
  };

  return (
    <div className="container mx-auto py-8 flex flex-col items-center justify-center min-h-screen">
      <form
        className="w-full max-w-lg mx-auto bg-white p-8 rounded-md shadow-md"
        onSubmit={handleLogin}
      >
        <div className="flex justify-center mb-6">
          <Image
            src={logo}
            alt="Panchang API Gateway Logo"
            width={150}
            height={150}
          />
        </div>
        <h1 className="font-bold mb-3 text-center text-xl">
          Panchang API Gateway Login
        </h1>
        <p className="text-center text-gray-600 mb-6">
          Please enter your username and password.
        </p>

        {error && <div className="text-red-500 text-sm mb-4">{error}</div>}

        <div className="mb-4">
          <Label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="username"
          >
            Username
          </Label>
          <input
            className="w-full px-4 py-2 border-gray-300 rounded-md focus:outline-none focus:border-blue-500"
            id="username"
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="Enter your username"
          />
        </div>

        <div className="mb-4">
          <Label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="password"
          >
            Password
          </Label>
          <input
            className="w-full px-4 py-2 border-gray-300 rounded-md focus:outline-none focus:border-blue-500"
            id="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Enter your password"
          />
        </div>

        <Button
          type="submit"
          color="orange"
          size="xs"
          className="w-full text-white bg-[#723B13] text-xl font-bold py-3 rounded-md transition duration-300"
        >
          Login
        </Button>
      </form>
    </div>
  );
}
