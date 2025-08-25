using Xunit;
using Microsoft.AspNetCore.TestHost;
using Microsoft.AspNetCore.Hosting;
using System.Net.Http;
using System.Threading.Tasks;

namespace dotnet_app.Tests;

public class HealthTest
{
    private readonly HttpClient _client;

    public HealthTest()
    {
        var builder = new WebHostBuilder()
            .UseStartup<Startup>();
        var server = new TestServer(builder);
        _client = server.CreateClient();
    }

    [Fact]
    public async Task GetHealthz_ReturnsOk()
    {
        var response = await _client.GetAsync("/healthz");
        response.EnsureSuccessStatusCode();
        var content = await response.Content.ReadAsStringAsync();
        Assert.Equal("OK", content);
    }
}