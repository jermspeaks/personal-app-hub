export interface AppStatus {
  name: string;
  description: string;
  online: boolean;
  lastUpdated: string; // ISO 8601 date string
  ports: number[];
}

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8518';

export async function fetchStatus(): Promise<AppStatus[]> {
  const response = await fetch(`${API_URL}/api/status`);
  if (!response.ok) {
    throw new Error(`Failed to fetch status: ${response.statusText}`);
  }
  return response.json();
}
