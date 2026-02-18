import { useEffect, useState } from 'react';
import { fetchStatus, type AppStatus } from '@/api/status';
import { AppCard } from '@/components/AppCard';

function App() {
  const [apps, setApps] = useState<AppStatus[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const loadStatus = async () => {
    try {
      setError(null);
      const statuses = await fetchStatus();
      setApps(statuses);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load app status');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadStatus();

    // Auto-refresh every 30 seconds
    const interval = setInterval(loadStatus, 30000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900 py-8 px-4">
      <div className="max-w-7xl mx-auto">
        <header className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-2">
            Personal Apps Hub
          </h1>
          <p className="text-gray-600 dark:text-gray-400">
            Monitor the status and last updated time for all personal applications
          </p>
        </header>

        {loading && (
          <div className="text-center py-12">
            <div className="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900 dark:border-gray-100"></div>
            <p className="mt-4 text-gray-600 dark:text-gray-400">Loading app status...</p>
          </div>
        )}

        {error && (
          <div className="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg p-4 mb-6">
            <p className="text-red-800 dark:text-red-200">
              Error: {error}
            </p>
            <button
              onClick={loadStatus}
              className="mt-2 px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition-colors"
            >
              Retry
            </button>
          </div>
        )}

        {!loading && !error && (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {apps.map((app) => (
              <AppCard key={app.name} app={app} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

export default App;
