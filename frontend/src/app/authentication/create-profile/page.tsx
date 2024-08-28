"use client";

import React, { useState } from "react";
import { Button, Label } from "flowbite-react";

export default function CreatePasswordPage() {
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");

  const handleCreatePassword = (e: any) => {
    e.preventDefault();
    setError("");

    // Client-side validation
    if (!password || !confirmPassword) {
      setError("Both fields are required.");
      return;
    }

    if (password !== confirmPassword) {
      setError("Passwords do not match.");
      return;
    }

    if (password.length < 8) {
      setError("Password must be at least 8 characters long.");
      return;
    }

    console.log("Password created successfully!");
    // Add logic to handle password creation
  };

  return (
    <div className="container mx-auto py-8 flex flex-col items-center justify-center min-h-screen">
      <form
        className="w-full max-w-lg mx-auto bg-white p-8 rounded-md shadow-md"
        onSubmit={handleCreatePassword}
      >
        <h1 className="font-bold mb-3 text-center text-xl">
          Create New Password
        </h1>
        <p className="text-center text-gray-600 mb-6">
          Please enter and confirm your new password.
        </p>

        {error && <div className="text-red-500 text-sm mb-4">{error}</div>}

        <div className="mb-4">
          <Label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="password"
          >
            New Password
          </Label>
          <input
            className="w-full px-4 py-2 border-gray-300 rounded-md focus:outline-none focus:border-blue-500"
            id="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Enter new password"
          />
        </div>

        <div className="mb-4">
          <Label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="confirmPassword"
          >
            Confirm Password
          </Label>
          <input
            className="w-full px-4 py-2 border-gray-300 rounded-md focus:outline-none focus:border-blue-500"
            id="confirmPassword"
            type="password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            placeholder="Confirm new password"
          />
        </div>

        <Button
          type="submit"
          color="orange"
          size="xs"
          className="w-full text-white bg-[#723B13] text-xl font-bold py-3 rounded-md transition duration-300"
        >
          Create Profile
        </Button>
      </form>
    </div>
  );
}
