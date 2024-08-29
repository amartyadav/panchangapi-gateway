import axios, { AxiosResponse } from "axios";

const BACKEND_URL = "http://localhost:1323";

interface UserRegistrationRequest {
  sessionToken: string;
  password: string;
}

interface UserRegistrationResponse {
  success?: string;
  api_key?: string;
  error?: string;
}

export const createProfile = async (
  data: UserRegistrationRequest
): Promise<UserRegistrationResponse> => {
  try {
    const response: AxiosResponse<UserRegistrationResponse> = await axios.post(
      `${BACKEND_URL}/createProfile`,
      data
    );
    return response.data;
  } catch (error: any) {
    if (error.response) {
      const errorMessage =
        error.response.data.error || "Unknown error occurred";
      console.error("Error creating profile:", errorMessage);
      throw new Error(errorMessage);
    } else {
      console.error("Error creating profile:", error.message);
      throw new Error("Network error");
    }
  }
};
