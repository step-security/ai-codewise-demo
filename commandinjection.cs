using System.Diagnostics;

namespace Injections
{
    public class OsCommandInjection
    {
        public void RunOsCommand(string command)
        {
            // ruleid: os-command-injection
            var process = Process.Start(command);
        }


        public void RunOsCommandWithProcessParam(string command)
        {
            Process process = new Process();

            process.StartInfo.FileName = command;
            // ruleid: os-command-injection
            process.Start();
        }

        public void RunConstantAppWithArgs(string args)
        {
            ProcessStartInfo processStartInfo = new ProcessStartInfo()
            {
                FileName = "constant",
                Arguments = "constant"
            };

            // ok: os-command-injection
            var process = Process.Start(processStartInfo);
        }
    }
}
