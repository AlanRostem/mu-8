const MAX_MEMORY_SIZE: i32 = 4096;

pub const Memory = struct {
    data: [MAX_MEMORY_SIZE]u8,
    pub fn new() Memory {
        return .{ .data = .{0} ** MAX_MEMORY_SIZE };
    }

    pub fn write(self: *Memory, address: u12, data: u8) void {
        self.data[address] = data;
    }

    pub fn read(self: *Memory, address: u12) u8 {
        return self.data[address];
    }
};
