import { AppStatus } from '@/api/status';

interface AppCardProps {
  app: AppStatus;
}

function formatTimeAgo(dateString: string): string {
  if (!dateString) {
    return 'Unknown';
  }

  const date = new Date(dateString);
  const now = new Date();
  const diffMs = now.getTime() - date.getTime();
  const diffSeconds = Math.floor(diffMs / 1000);
  const diffMinutes = Math.floor(diffSeconds / 60);
  const diffHours = Math.floor(diffMinutes / 60);
  const diffDays = Math.floor(diffHours / 24);

  if (diffDays > 0) {
    return `${diffDays} day${diffDays === 1 ? '' : 's'} ago`;
  }
  if (diffHours > 0) {
    return `${diffHours} hour${diffHours === 1 ? '' : 's'} ago`;
  }
  if (diffMinutes > 0) {
    return `${diffMinutes} minute${diffMinutes === 1 ? '' : 's'} ago`;
  }
  return 'Just now';
}

export function AppCard({ app }: AppCardProps) {
  return (
    <div className="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6 border border-gray-200 dark:border-gray-700 hover:shadow-lg transition-shadow">
      <div className="flex items-start justify-between mb-4">
        <h2 className="text-xl font-semibold text-gray-900 dark:text-gray-100">
          {app.name}
        </h2>
        <span
          className={`px-3 py-1 rounded-full text-sm font-medium ${
            app.online
              ? 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200'
              : 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200'
          }`}
        >
          {app.online ? 'Online' : 'Offline'}
        </span>
      </div>
      <p className="text-gray-600 dark:text-gray-400 mb-4 text-sm">
        {app.description}
      </p>
      <div className="flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
        <span>Last updated: {formatTimeAgo(app.lastUpdated)}</span>
        <span className="text-xs">Ports: {app.ports.join(', ')}</span>
      </div>
    </div>
  );
}
