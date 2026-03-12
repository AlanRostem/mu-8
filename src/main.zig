const std = @import("std");
const memory = @import("memory.zig");
pub fn main() !void {
    var mem = memory.Memory.new();
    mem.write(0, 123);
    // Prints to stderr, ignoring potential errors.
    std.debug.print("Memory written: {d}!\n", .{mem.read(0)});
}
